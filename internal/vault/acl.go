package vault

// Permission represents an allowed action within a vault.
type Permission string

const (
	PermRead   Permission = "read"
	PermWrite  Permission = "write"
	PermDelete Permission = "delete"
	PermAdmin  Permission = "admin"
)

// ACL manages the permissions for a specific resource.
type ACL struct {
	ResourceID  string
	Permissions map[string][]Permission // Key: UserID or PeerID
}

// NewACL creates a new ACL for a resource.
func NewACL(resourceID string) *ACL {
	return &ACL{
		ResourceID:  resourceID,
		Permissions: make(map[string][]Permission),
	}
}

// Grant adds a permission for a user.
func (a *ACL) Grant(userID string, perm Permission) {
	a.Permissions[userID] = append(a.Permissions[userID], perm)
}

// HasPermission checks if a user has a specific permission.
func (a *ACL) HasPermission(userID string, perm Permission) bool {
	perms, exists := a.Permissions[userID]
	if !exists {
		return false
	}
	for _, p := range perms {
		if p == perm || p == PermAdmin {
			return true
		}
	}
	return false
}
