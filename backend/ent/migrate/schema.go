// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BloodtypesColumns holds the columns for the "bloodtypes" table.
	BloodtypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "btname", Type: field.TypeString},
	}
	// BloodtypesTable holds the schema information for the "bloodtypes" table.
	BloodtypesTable = &schema.Table{
		Name:        "bloodtypes",
		Columns:     BloodtypesColumns,
		PrimaryKey:  []*schema.Column{BloodtypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CertificatesColumns holds the columns for the "certificates" table.
	CertificatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "certificate_name", Type: field.TypeString},
	}
	// CertificatesTable holds the schema information for the "certificates" table.
	CertificatesTable = &schema.Table{
		Name:        "certificates",
		Columns:     CertificatesColumns,
		PrimaryKey:  []*schema.Column{CertificatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CoveredPersonsColumns holds the columns for the "covered_persons" table.
	CoveredPersonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "Certificate_id", Type: field.TypeInt, Nullable: true},
		{Name: "Fund_id", Type: field.TypeInt, Nullable: true},
		{Name: "Patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "SchemeType_id", Type: field.TypeInt, Nullable: true},
	}
	// CoveredPersonsTable holds the schema information for the "covered_persons" table.
	CoveredPersonsTable = &schema.Table{
		Name:       "covered_persons",
		Columns:    CoveredPersonsColumns,
		PrimaryKey: []*schema.Column{CoveredPersonsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "covered_persons_certificates_Certificate_CoveredPerson",
				Columns: []*schema.Column{CoveredPersonsColumns[1]},

				RefColumns: []*schema.Column{CertificatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "covered_persons_funds_Fund_CoveredPerson",
				Columns: []*schema.Column{CoveredPersonsColumns[2]},

				RefColumns: []*schema.Column{FundsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "covered_persons_patients_Patient_CoveredPerson",
				Columns: []*schema.Column{CoveredPersonsColumns[3]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "covered_persons_scheme_types_SchemeType_CoveredPerson",
				Columns: []*schema.Column{CoveredPersonsColumns[4]},

				RefColumns: []*schema.Column{SchemeTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "department_name", Type: field.TypeString},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DiagnosesColumns holds the columns for the "diagnoses" table.
	DiagnosesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "department_id", Type: field.TypeInt, Nullable: true},
		{Name: "disease_id", Type: field.TypeInt, Nullable: true},
		{Name: "doctor_id", Type: field.TypeInt, Nullable: true},
		{Name: "patient_id", Type: field.TypeInt, Nullable: true},
	}
	// DiagnosesTable holds the schema information for the "diagnoses" table.
	DiagnosesTable = &schema.Table{
		Name:       "diagnoses",
		Columns:    DiagnosesColumns,
		PrimaryKey: []*schema.Column{DiagnosesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "diagnoses_departments_department_diagnose",
				Columns: []*schema.Column{DiagnosesColumns[1]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "diagnoses_diseases_disease_diagnose",
				Columns: []*schema.Column{DiagnosesColumns[2]},

				RefColumns: []*schema.Column{DiseasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "diagnoses_doctors_doctor_diagnose",
				Columns: []*schema.Column{DiagnosesColumns[3]},

				RefColumns: []*schema.Column{DoctorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "diagnoses_patients_patient_diagnose",
				Columns: []*schema.Column{DiagnosesColumns[4]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DiseasesColumns holds the columns for the "diseases" table.
	DiseasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "disease_name", Type: field.TypeString},
	}
	// DiseasesTable holds the schema information for the "diseases" table.
	DiseasesTable = &schema.Table{
		Name:        "diseases",
		Columns:     DiseasesColumns,
		PrimaryKey:  []*schema.Column{DiseasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DoctorsColumns holds the columns for the "doctors" table.
	DoctorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "doctor_name", Type: field.TypeString},
		{Name: "doctor_password", Type: field.TypeString},
		{Name: "doctor_email", Type: field.TypeString},
		{Name: "doctor_tel", Type: field.TypeString},
	}
	// DoctorsTable holds the schema information for the "doctors" table.
	DoctorsTable = &schema.Table{
		Name:        "doctors",
		Columns:     DoctorsColumns,
		PrimaryKey:  []*schema.Column{DoctorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FundsColumns holds the columns for the "funds" table.
	FundsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fund_name", Type: field.TypeString},
	}
	// FundsTable holds the schema information for the "funds" table.
	FundsTable = &schema.Table{
		Name:        "funds",
		Columns:     FundsColumns,
		PrimaryKey:  []*schema.Column{FundsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gname", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MedicalsColumns holds the columns for the "medicals" table.
	MedicalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "medical_name", Type: field.TypeString},
		{Name: "medical_email", Type: field.TypeString},
		{Name: "medical_password", Type: field.TypeString},
		{Name: "medical_tel", Type: field.TypeString},
	}
	// MedicalsTable holds the schema information for the "medicals" table.
	MedicalsTable = &schema.Table{
		Name:        "medicals",
		Columns:     MedicalsColumns,
		PrimaryKey:  []*schema.Column{MedicalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// NursesColumns holds the columns for the "nurses" table.
	NursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "nurse_name", Type: field.TypeString},
		{Name: "nurse_email", Type: field.TypeString},
		{Name: "nurse_password", Type: field.TypeString},
		{Name: "nurse_tel", Type: field.TypeString},
	}
	// NursesTable holds the schema information for the "nurses" table.
	NursesTable = &schema.Table{
		Name:        "nurses",
		Columns:     NursesColumns,
		PrimaryKey:  []*schema.Column{NursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "patient_name", Type: field.TypeString},
		{Name: "patient_age", Type: field.TypeInt},
		{Name: "patient_weight", Type: field.TypeFloat64},
		{Name: "patient_height", Type: field.TypeFloat64},
		{Name: "bloodtype_id", Type: field.TypeInt, Nullable: true},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "prefix_id", Type: field.TypeInt, Nullable: true},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:       "patients",
		Columns:    PatientsColumns,
		PrimaryKey: []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "patients_bloodtypes_frombloodtype",
				Columns: []*schema.Column{PatientsColumns[5]},

				RefColumns: []*schema.Column{BloodtypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_genders_fromgender",
				Columns: []*schema.Column{PatientsColumns[6]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patients_prefixes_fromprefix",
				Columns: []*schema.Column{PatientsColumns[7]},

				RefColumns: []*schema.Column{PrefixesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PrefixesColumns holds the columns for the "prefixes" table.
	PrefixesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "pname", Type: field.TypeString},
	}
	// PrefixesTable holds the schema information for the "prefixes" table.
	PrefixesTable = &schema.Table{
		Name:        "prefixes",
		Columns:     PrefixesColumns,
		PrimaryKey:  []*schema.Column{PrefixesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RentsColumns holds the columns for the "rents" table.
	RentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "nurse_id", Type: field.TypeInt, Nullable: true},
		{Name: "Patient_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "room_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// RentsTable holds the schema information for the "rents" table.
	RentsTable = &schema.Table{
		Name:       "rents",
		Columns:    RentsColumns,
		PrimaryKey: []*schema.Column{RentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "rents_nurses_fromnurse",
				Columns: []*schema.Column{RentsColumns[2]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "rents_patients_frompatient",
				Columns: []*schema.Column{RentsColumns[3]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "rents_rooms_rents",
				Columns: []*schema.Column{RentsColumns[4]},

				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "building", Type: field.TypeInt},
		{Name: "floor", Type: field.TypeInt},
		{Name: "roomtype_id", Type: field.TypeInt, Nullable: true},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "rooms_roomtypes_rooms",
				Columns: []*schema.Column{RoomsColumns[4]},

				RefColumns: []*schema.Column{RoomtypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RoomtypesColumns holds the columns for the "roomtypes" table.
	RoomtypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "bathroom", Type: field.TypeInt},
		{Name: "toilets", Type: field.TypeInt},
		{Name: "areasize", Type: field.TypeFloat64},
		{Name: "etc", Type: field.TypeString},
	}
	// RoomtypesTable holds the schema information for the "roomtypes" table.
	RoomtypesTable = &schema.Table{
		Name:        "roomtypes",
		Columns:     RoomtypesColumns,
		PrimaryKey:  []*schema.Column{RoomtypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SchemeTypesColumns holds the columns for the "scheme_types" table.
	SchemeTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "scheme_type_name", Type: field.TypeString},
	}
	// SchemeTypesTable holds the schema information for the "scheme_types" table.
	SchemeTypesTable = &schema.Table{
		Name:        "scheme_types",
		Columns:     SchemeTypesColumns,
		PrimaryKey:  []*schema.Column{SchemeTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BloodtypesTable,
		CertificatesTable,
		CoveredPersonsTable,
		DepartmentsTable,
		DiagnosesTable,
		DiseasesTable,
		DoctorsTable,
		FundsTable,
		GendersTable,
		MedicalsTable,
		NursesTable,
		PatientsTable,
		PrefixesTable,
		RentsTable,
		RoomsTable,
		RoomtypesTable,
		SchemeTypesTable,
	}
)

func init() {
	CoveredPersonsTable.ForeignKeys[0].RefTable = CertificatesTable
	CoveredPersonsTable.ForeignKeys[1].RefTable = FundsTable
	CoveredPersonsTable.ForeignKeys[2].RefTable = PatientsTable
	CoveredPersonsTable.ForeignKeys[3].RefTable = SchemeTypesTable
	DiagnosesTable.ForeignKeys[0].RefTable = DepartmentsTable
	DiagnosesTable.ForeignKeys[1].RefTable = DiseasesTable
	DiagnosesTable.ForeignKeys[2].RefTable = DoctorsTable
	DiagnosesTable.ForeignKeys[3].RefTable = PatientsTable
	PatientsTable.ForeignKeys[0].RefTable = BloodtypesTable
	PatientsTable.ForeignKeys[1].RefTable = GendersTable
	PatientsTable.ForeignKeys[2].RefTable = PrefixesTable
	RentsTable.ForeignKeys[0].RefTable = NursesTable
	RentsTable.ForeignKeys[1].RefTable = PatientsTable
	RentsTable.ForeignKeys[2].RefTable = RoomsTable
	RoomsTable.ForeignKeys[0].RefTable = RoomtypesTable
}
