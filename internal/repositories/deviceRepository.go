package repositories

import (
	"database/sql"
	"fmt"

	"github.com/entraktest/configs"
	"github.com/entraktest/internal/models"
	"github.com/entraktest/pkg/mysql"
)

type DeviceRepository struct {
	db     mysql.MySqlFactory
	config *configs.Config
}

type IDeviceRepository interface {
	Store(deviceRequest models.Device) error
	GetById(id uint64) (models.Device, error)
	GetAll() ([]models.Device, error)
}

func NewDeviceRepository(mysqlDB mysql.MySqlFactory, cfg *configs.Config) *DeviceRepository {
	return &DeviceRepository{
		db:     mysqlDB,
		config: cfg,
	}
}

func (d *DeviceRepository) GetAll() (devices []models.Device, err error) {
	db, err := d.db.GetDB()
	if err != nil {
		return
	}

	rows, err := db.Query(fmt.Sprintf(`select id, device, unit, value, datetime from deviceMaster`))

	if err != nil {
		return
	}

	for rows.Next() {
		var device models.Device
		err = rows.Scan(
			&device.ID,
			&device.Device,
			&device.Unit,
			&device.Value,
			&device.Datetime,
		)

		devices = append(devices, device)
	}

	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func (d *DeviceRepository) GetById(id uint64) (device models.Device, err error) {
	db, err := d.db.GetDB()
	if err != nil {
		return
	}

	rows, err := db.Query(fmt.Sprintf(`
		select id, device, unit, value, datetime from deviceMaster
		WHERE id=?
	`), id)

	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(
			&device.ID,
			&device.Device,
			&device.Unit,
			&device.Value,
			&device.Datetime,
		)
	}

	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func (d *DeviceRepository) Store(device models.Device) (err error) {
	db, err := d.db.GetDB()
	if err != nil {
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	stmt, err := tx.Prepare("INSERT INTO deviceMaster (device, unit, value, datetime) VALUES (?,?,?,?)")
	if err != nil {
		return
	}
	res, err := stmt.Exec(device.Device, device.Unit, device.Value, device.Datetime)
	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return errTx
		}
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return errTx
		}
		return
	}

	err = tx.Commit()

	return
}
