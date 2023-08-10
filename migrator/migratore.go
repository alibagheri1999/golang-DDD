package migrator

import (
	"embed"
	"errors"
	"log"
	"remote-task/infrastructure/persistence/mysql"
	"time"

	migrate "github.com/rubenv/sql-migrate"
)

const (
	MigrationsTable = "migrator"
)

//go:embed migrations/*.sql
var fsMigrations embed.FS

func New(repo *mysql.Repositories) *Migrator {
	return &Migrator{
		repo: repo,
	}
}

type Migrator struct {
	repo *mysql.Repositories
}

func (m *Migrator) Run(action string, count int) error {

	if action == "new" {
		log.Println(time.Now().Unix())
		return nil
	}

	if action == "fresh" {
		if err := m.dropAllTables(); err != nil {
			log.Fatalf("dropping all tables: %v", err)
		} else {
			log.Println("all tables dropped")
		}

		// after dropping tables we will run all the migrator
		action = "up"
		// zero count for the migrator agent means to apply/rollback all residues
		count = 0
	}
	plannedItems, err := m.getFilesSource().FindMigrations()
	if err != nil {
		log.Fatalf("getting list of available migrator: %v", err)
	}

	appliedItems, err := m.getAlreadyApplied()
	if err != nil {
		log.Fatalf("getting list of applied migrator: %v", err)
	}

	if action == "up" {
		if err := m.applyMigration(count, plannedItems, appliedItems); err != nil {
			return err
		}
	}

	if action == "down" {
		if err := m.undoMigration(count, appliedItems); err != nil {
			return err
		}
	}

	appliedItems, err = m.getAlreadyApplied()
	if err != nil {
		log.Fatalf("getting list of applied migrator: %v", err)
	}

	m.reportStatus(plannedItems, appliedItems)

	return nil
}

func (m *Migrator) agent() migrate.MigrationSet {
	return migrate.MigrationSet{
		TableName:     MigrationsTable,
		IgnoreUnknown: false,
	}
}

func (m *Migrator) getFilesSource() *migrate.EmbedFileSystemMigrationSource {
	return &migrate.EmbedFileSystemMigrationSource{
		FileSystem: fsMigrations,
		Root:       "migrator",
	}
}

func (m *Migrator) dropAllTables() error {

	tx, err := m.repo.DB().Begin()
	if err != nil {
		log.Fatalf("creating transaction: %v", err)
	}

	rows, err := tx.Query(
		"SELECT concat('DROP TABLE IF EXISTS `', table_name, '`;') AS `stmt` FROM information_schema.tables WHERE table_schema = ?;",
		"mysql",
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	stmts := make([]string, 0)
	for rows.Next() {
		var stmt string
		if err = rows.Scan(&stmt); err != nil {
			_ = tx.Rollback()
			return err
		}
		stmts = append(stmts, stmt)
	}
	if err := rows.Close(); err != nil {
		return err
	}

	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS=0;"); err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS=1;"); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (m *Migrator) getAlreadyApplied() ([]*migrate.MigrationRecord, error) {
	return m.agent().GetMigrationRecords(m.repo.DB(), mysql.Dialect)
}

func (m *Migrator) applyMigration(count int, plannedItems []*migrate.Migration, appliedItems []*migrate.MigrationRecord) error {
	if len(plannedItems) == len(appliedItems) {
		log.Println("no new migration available to apply")
		return nil
	}
	affected, err := m.agent().ExecMax(m.repo.DB(), mysql.Dialect, m.getFilesSource(), migrate.Up, count)
	if err != nil {
		return err
	}
	log.Printf("%v migrator applied\n", affected)
	return nil
}

func (m *Migrator) undoMigration(count int, appliedItems []*migrate.MigrationRecord) error {
	if len(appliedItems) == 0 {
		return errors.New("no applied migration available to roll back")
	}
	affected, err := m.agent().ExecMax(m.repo.DB(), mysql.Dialect, m.getFilesSource(), migrate.Down, count)
	if err != nil {
		return err
	}
	log.Printf("%v migrator rolled back\n", affected)
	return nil
}

func (m *Migrator) reportStatus(plannedItems []*migrate.Migration, appliedItems []*migrate.MigrationRecord) {
	for _, item := range plannedItems {
		var appliedAt time.Time
		for _, record := range appliedItems {
			if record.Id == item.Id {
				appliedAt = record.AppliedAt
				break
			}
		}
		if appliedAt.IsZero() {
			log.Printf("%s - Not Applied\n", item.Id)
		} else {
			log.Printf("%s - Applied at %s\n", item.Id, appliedAt)
		}
	}
}
