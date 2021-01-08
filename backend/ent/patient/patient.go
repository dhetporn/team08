// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPatientName holds the string denoting the patient_name field in the database.
	FieldPatientName = "patient_name"
	// FieldPatientAge holds the string denoting the patient_age field in the database.
	FieldPatientAge = "patient_age"
	// FieldPatientWeight holds the string denoting the patient_weight field in the database.
	FieldPatientWeight = "patient_weight"
	// FieldPatientHeight holds the string denoting the patient_height field in the database.
	FieldPatientHeight = "patient_height"

	// EdgeFrompatient holds the string denoting the frompatient edge name in mutations.
	EdgeFrompatient = "frompatient"
	// EdgePatientCoveredPerson holds the string denoting the patient_coveredperson edge name in mutations.
	EdgePatientCoveredPerson = "Patient_CoveredPerson"
	// EdgePatientDiagnose holds the string denoting the patient_diagnose edge name in mutations.
	EdgePatientDiagnose = "patient_diagnose"
	// EdgeGender holds the string denoting the gender edge name in mutations.
	EdgeGender = "gender"
	// EdgePrefix holds the string denoting the prefix edge name in mutations.
	EdgePrefix = "prefix"
	// EdgeBloodtype holds the string denoting the bloodtype edge name in mutations.
	EdgeBloodtype = "bloodtype"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// FrompatientTable is the table the holds the frompatient relation/edge.
	FrompatientTable = "rents"
	// FrompatientInverseTable is the table name for the Rent entity.
	// It exists in this package in order to avoid circular dependency with the "rent" package.
	FrompatientInverseTable = "rents"
	// FrompatientColumn is the table column denoting the frompatient relation/edge.
	FrompatientColumn = "Patient_id"
	// PatientCoveredPersonTable is the table the holds the Patient_CoveredPerson relation/edge.
	PatientCoveredPersonTable = "covered_persons"
	// PatientCoveredPersonInverseTable is the table name for the CoveredPerson entity.
	// It exists in this package in order to avoid circular dependency with the "coveredperson" package.
	PatientCoveredPersonInverseTable = "covered_persons"
	// PatientCoveredPersonColumn is the table column denoting the Patient_CoveredPerson relation/edge.
	PatientCoveredPersonColumn = "Patient_id"
	// PatientDiagnoseTable is the table the holds the patient_diagnose relation/edge.
	PatientDiagnoseTable = "diagnoses"
	// PatientDiagnoseInverseTable is the table name for the Diagnose entity.
	// It exists in this package in order to avoid circular dependency with the "diagnose" package.
	PatientDiagnoseInverseTable = "diagnoses"
	// PatientDiagnoseColumn is the table column denoting the patient_diagnose relation/edge.
	PatientDiagnoseColumn = "patient_id"
	// GenderTable is the table the holds the gender relation/edge.
	GenderTable = "patients"
	// GenderInverseTable is the table name for the Gender entity.
	// It exists in this package in order to avoid circular dependency with the "gender" package.
	GenderInverseTable = "genders"
	// GenderColumn is the table column denoting the gender relation/edge.
	GenderColumn = "gender_id"
	// PrefixTable is the table the holds the prefix relation/edge.
	PrefixTable = "patients"
	// PrefixInverseTable is the table name for the Prefix entity.
	// It exists in this package in order to avoid circular dependency with the "prefix" package.
	PrefixInverseTable = "prefixes"
	// PrefixColumn is the table column denoting the prefix relation/edge.
	PrefixColumn = "prefix_id"
	// BloodtypeTable is the table the holds the bloodtype relation/edge.
	BloodtypeTable = "patients"
	// BloodtypeInverseTable is the table name for the Bloodtype entity.
	// It exists in this package in order to avoid circular dependency with the "bloodtype" package.
	BloodtypeInverseTable = "bloodtypes"
	// BloodtypeColumn is the table column denoting the bloodtype relation/edge.
	BloodtypeColumn = "bloodtype_id"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldPatientName,
	FieldPatientAge,
	FieldPatientWeight,
	FieldPatientHeight,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Patient type.
var ForeignKeys = []string{
	"bloodtype_id",
	"gender_id",
	"prefix_id",
}

var (
	// PatientNameValidator is a validator for the "Patient_name" field. It is called by the builders before save.
	PatientNameValidator func(string) error
	// PatientAgeValidator is a validator for the "Patient_age" field. It is called by the builders before save.
	PatientAgeValidator func(int) error
	// PatientWeightValidator is a validator for the "Patient_weight" field. It is called by the builders before save.
	PatientWeightValidator func(float64) error
	// PatientHeightValidator is a validator for the "Patient_height" field. It is called by the builders before save.
	PatientHeightValidator func(float64) error
)
