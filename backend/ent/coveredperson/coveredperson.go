// Code generated by entc, DO NOT EDIT.

package coveredperson

const (
	// Label holds the string label denoting the coveredperson type in the database.
	Label = "covered_person"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"

	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "Patient"
	// EdgeSchemeType holds the string denoting the schemetype edge name in mutations.
	EdgeSchemeType = "SchemeType"
	// EdgeFund holds the string denoting the fund edge name in mutations.
	EdgeFund = "Fund"
	// EdgeCertificate holds the string denoting the certificate edge name in mutations.
	EdgeCertificate = "Certificate"

	// Table holds the table name of the coveredperson in the database.
	Table = "covered_persons"
	// PatientTable is the table the holds the Patient relation/edge.
	PatientTable = "covered_persons"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the Patient relation/edge.
	PatientColumn = "Patient_id"
	// SchemeTypeTable is the table the holds the SchemeType relation/edge.
	SchemeTypeTable = "covered_persons"
	// SchemeTypeInverseTable is the table name for the SchemeType entity.
	// It exists in this package in order to avoid circular dependency with the "schemetype" package.
	SchemeTypeInverseTable = "scheme_types"
	// SchemeTypeColumn is the table column denoting the SchemeType relation/edge.
	SchemeTypeColumn = "SchemeType_id"
	// FundTable is the table the holds the Fund relation/edge.
	FundTable = "covered_persons"
	// FundInverseTable is the table name for the Fund entity.
	// It exists in this package in order to avoid circular dependency with the "fund" package.
	FundInverseTable = "funds"
	// FundColumn is the table column denoting the Fund relation/edge.
	FundColumn = "Fund_id"
	// CertificateTable is the table the holds the Certificate relation/edge.
	CertificateTable = "covered_persons"
	// CertificateInverseTable is the table name for the Certificate entity.
	// It exists in this package in order to avoid circular dependency with the "certificate" package.
	CertificateInverseTable = "certificates"
	// CertificateColumn is the table column denoting the Certificate relation/edge.
	CertificateColumn = "Certificate_id"
)

// Columns holds all SQL columns for coveredperson fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the CoveredPerson type.
var ForeignKeys = []string{
	"Certificate_id",
	"Fund_id",
	"Patient_id",
	"SchemeType_id",
}