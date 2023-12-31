package types

type AuthRole string

const (
	AuthRoleVerified AuthRole = "Verified"
	AuthRoleMember   AuthRole = "Member"
	AuthRoleElder    AuthRole = "Elder"
	AuthRoleCoLeader AuthRole = "Co-Leader"
	AuthRoleLeader   AuthRole = "Leader"
	AuthRoleAdmin    AuthRole = "~~~admin~~~"
)

func (r AuthRole) String() string {
	switch r {
	case AuthRoleVerified:
		return "Verifiziert"
	case AuthRoleMember:
		return "Mitglied"
	case AuthRoleElder:
		return "Ältester"
	case AuthRoleCoLeader:
		return "Vize-Anführer"
	case AuthRoleLeader:
		return "Anführer"
	case AuthRoleAdmin:
		return "Administrator"
	default:
		return "Ungültige Rolle"
	}
}
