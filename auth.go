package main

import (
	"strings"

	"gopkg.in/kataras/iris.v6"
)

func loginRequired(ctx *iris.Context) {
	if !isSignedIn(ctx) {
		redirect := conf.AppUrl + "/sign-in"
		path := ctx.Request.URL.String()
		if len(path) > 0 {
			redirect += "?next=" + path
		}
		ctx.Redirect(redirect)
	} else {
		if conf.DemoLock > 0 {
			session := ctx.Session()
			userId := session.Get(USER_ID)
			if userId.(uint) == conf.DemoLock {
				forbiddenGet := []string{appPath() + "/account/delete", appPath() + "/websites/delete/", appPath() + "/websites/default/"}
				forbiddenPost := []string{appPath() + "/websites/new", appPath() + "/websites/", appPath() + "/account", appPath() + "/account/settings"}
				session := ctx.Session()

				user := &User{}
				db.First(user, userId.(uint))

				if ctx.Request.Method == "GET" {
					for _, fg := range forbiddenGet {
						if strings.HasPrefix(ctx.Request.URL.String(), fg) {
							session.SetFlash("error", "This is a demo account and you are not allowed to do that action.")
							user.redirectToDefaultWebsite(ctx)
						}
					}
				}

				if ctx.Request.Method == "POST" {
					for _, fp := range forbiddenPost {
						if strings.HasPrefix(ctx.Request.URL.String(), fp) || ctx.Request.URL.String() == fp {
							session.SetFlash("error", "This is a demo account and you are not allowed to do that action.")
							user.redirectToDefaultWebsite(ctx)
						}
					}
				}
			}
		}
		ctx.Next()
	}
}

func signIn(ctx *iris.Context, user *User) {
	session := ctx.Session()
	session.Set(USER_ID, user.ID)
}

func signOut(ctx *iris.Context) {
	session := ctx.Session()
	session.Delete(USER_ID)
}

func isSignedIn(ctx *iris.Context) bool {
	session := ctx.Session()
	userId := session.Get(USER_ID)
	if userId == nil {
		return false
	}
	if userId.(uint) > 0 {
		u := &User{}
		db.First(u, userId.(uint))
		return u.ID > 0 && u.Verified
	}
	return false
}

func getSignedInUserId(ctx *iris.Context) uint {
	session := ctx.Session()
	userId := session.Get(USER_ID)
	if userId == nil {
		return uint(0)
	} else {
		return userId.(uint)
	}
}
