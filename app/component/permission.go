package component

import "log"

func CheckPermission(account string, permission string, method string) bool {
	if account == "admin" {
		return true
	}
	return hasPermissionForUser(account, permission, method)
}

// 检查是否有这个权限
func hasPermissionForUser(account string, permission string, method string) bool {
	roles, _ := enforcer.GetRolesForUser(account)
	for _, role := range roles {
		ok := enforcer.HasPermissionForUser(role, permission, method)
		if ok {
			return ok
		}
	}
	return false
}

// 授权用户到角色
func AddRoleForUser(role string, account string) bool {
	ok, err := enforcer.AddRoleForUser(account, role)
	if err != nil {
		log.Fatal("授权用户到角色")
	}
	return ok
}

// 授权用户到角色 批量
func AddRolesForUser(account string, role []string) bool {
	ok, err := enforcer.AddRolesForUser(account, role)
	if err != nil {
		log.Fatal("授权用户到角色")
	}
	return ok
}

//  添加权限到角色
func AddPermissionForUser(permission string, method string, role string) bool {
	ok, err := enforcer.AddPermissionForUser(role, permission, method)
	if err != nil {
		log.Fatal("添加权限到角色")
	}
	return ok
}

// 获取用户角色
func GetRolesForUser(user string) []string {
	roles, err := enforcer.GetRolesForUser(user)
	if err != nil {
		log.Fatal("获取用户角色")
	}
	return roles
}

// 获取用户（角色）权限
func GetPermissionsForRole(role string) []map[string]string {
	var permissions []map[string]string
	currentPermissions := enforcer.GetPermissionsForUser(role)
	for _, currentPermission := range currentPermissions {
		permissions = append(permissions, map[string]string{
			"method":     currentPermission[2],
			"permission": currentPermission[1],
		})
	}
	return permissions
}

// 获取用户权限
func GetPermissionsForUser(account string) []map[string]string {
	var permissions []map[string]string
	roles := GetRolesForUser(account)
	for _, role := range roles {
		rolePermissions := GetPermissionsForRole(role)
		for _, rolePermission := range rolePermissions {
			permissions = append(permissions, rolePermission)
		}
	}
	return permissions
}

// 删除用户的所有角色
func DeleteRolesForUser(account string) bool {
	_, err := enforcer.DeleteRolesForUser(account)
	if err != nil {
		log.Fatal("删除用户的所有角色")
	}
	return true
}

// 删除角色的权限
func DeletePermissionsForUser(role string) bool {
	_, err := enforcer.DeletePermissionsForUser(role)
	if err != nil {
		log.Fatal("删除角色权限")
	}
	return true
}

// 删除拥有对应角色的(用户角色权限)
func DeleteRoleForUsers(role string) bool {
	users, err := enforcer.GetUsersForRole(role)
	if err != nil {
		log.Fatal("获取具有角色的用户")
	}
	for _, user := range users {
		_, _ = enforcer.DeleteRoleForUser(user, role)
	}
	return true
}
