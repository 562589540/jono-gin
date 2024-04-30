package temp

import (
	"github.com/562589540/jono-gin/ghub"
	"github.com/562589540/jono-gin/internal/app/system/model"
	"github.com/562589540/jono-gin/internal/constants"
	"github.com/goccy/go-json"
	"log"
)

type MenuInput struct {
	OriginalID      uint   `json:"id"`
	ParentID        uint   `json:"parentId"`
	MenuType        int    `json:"menuType"`
	Title           string `json:"title"`
	Name            string `json:"name"`
	Path            string `json:"path"`
	Component       string `json:"component"`
	Rank            int    `json:"rank"`
	Redirect        string `json:"redirect"`
	Icon            string `json:"icon"`
	ExtraIcon       string `json:"extraIcon"`
	EnterTransition string `json:"enterTransition"`
	LeaveTransition string `json:"leaveTransition"`
	ActivePath      string `json:"activePath"`
	Auths           string `json:"auths"`
	FrameSrc        string `json:"frameSrc"`
	FrameLoading    bool   `json:"frameLoading"`
	KeepAlive       bool   `json:"keepAlive"`
	HiddenTag       bool   `json:"hiddenTag"`
	FixedTag        bool   `json:"fixedTag"`
	ShowLink        bool   `json:"showLink"`
	ShowParent      bool   `json:"showParent"`
}

var jsonData = `{
  "success": true,
  "data": [
    {
      "parentId": 0,
      "id": 100,
      "menuType": 0,
      "title": "menus.pureExternalPage",
      "name": "PureIframe",
      "path": "/iframe",
      "component": "",
      "rank": 7,
      "redirect": "",
      "icon": "ri:links-fill",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 100,
      "id": 101,
      "menuType": 0,
      "title": "menus.pureExternalDoc",
      "name": "PureIframeExternal",
      "path": "/iframe/external",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 101,
      "id": 102,
      "menuType": 2,
      "title": "menus.pureExternalLink",
      "name": "https://yiming_chang.gitee.io/pure-admin-doc",
      "path": "/external",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 101,
      "id": 103,
      "menuType": 2,
      "title": "menus.pureUtilsLink",
      "name": "https://pure-admin-gutils.netlify.app/",
      "path": "/pureUtilsLink",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 100,
      "id": 104,
      "menuType": 1,
      "title": "menus.pureEmbeddedDoc",
      "name": "PureIframeEmbedded",
      "path": "/iframe/embedded",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 104,
      "id": 105,
      "menuType": 1,
      "title": "menus.pureEpDoc",
      "name": "FrameEp",
      "path": "/iframe/ep",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "https://element-plus.org/zh-CN/",
      "frameLoading": true,
      "keepAlive": true,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 104,
      "id": 106,
      "menuType": 1,
      "title": "menus.pureTailwindcssDoc",
      "name": "FrameTailwindcss",
      "path": "/iframe/tailwindcss",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "https://tailwindcss.com/docs/installation",
      "frameLoading": true,
      "keepAlive": true,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 104,
      "id": 107,
      "menuType": 1,
      "title": "menus.pureVueDoc",
      "name": "FrameVue",
      "path": "/iframe/vue3",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "https://cn.vuejs.org/",
      "frameLoading": true,
      "keepAlive": true,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 104,
      "id": 108,
      "menuType": 1,
      "title": "menus.pureViteDoc",
      "name": "FrameVite",
      "path": "/iframe/vite",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "https://cn.vitejs.dev/",
      "frameLoading": true,
      "keepAlive": true,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 104,
      "id": 109,
      "menuType": 1,
      "title": "menus.purePiniaDoc",
      "name": "FramePinia",
      "path": "/iframe/pinia",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "https://pinia.vuejs.org/zh/index.html",
      "frameLoading": true,
      "keepAlive": true,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 104,
      "id": 110,
      "menuType": 1,
      "title": "menus.pureRouterDoc",
      "name": "FrameRouter",
      "path": "/iframe/vue-router",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "https://router.vuejs.org/zh/",
      "frameLoading": true,
      "keepAlive": true,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 0,
      "id": 200,
      "menuType": 0,
      "title": "menus.purePermission",
      "name": "PurePermission",
      "path": "/permission",
      "component": "",
      "rank": 9,
      "redirect": "",
      "icon": "ep:lollipop",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 200,
      "id": 201,
      "menuType": 0,
      "title": "menus.purePermissionPage",
      "name": "PermissionPage",
      "path": "/permission/page/index",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 200,
      "id": 202,
      "menuType": 0,
      "title": "menus.purePermissionButton",
      "name": "PermissionButton",
      "path": "/permission/button/index",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 202,
      "id": 203,
      "menuType": 3,
      "title": "添加",
      "name": "2",
      "path": "",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "permission:btn:add",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 202,
      "id": 204,
      "menuType": 3,
      "title": "修改",
      "name": "3",
      "path": "",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "permission:btn:edit",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 202,
      "id": 205,
      "menuType": 3,
      "title": "删除",
      "name": "1",
      "path": "",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "permission:btn:delete",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 0,
      "id": 300,
      "menuType": 0,
      "title": "menus.pureSysManagement",
      "name": "PureSystem",
      "path": "/system",
      "component": "",
      "rank": 10,
      "redirect": "",
      "icon": "ri:settings-3-line",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 300,
      "id": 301,
      "menuType": 0,
      "title": "menus.pureUser",
      "name": "SystemUser",
      "path": "/system/user/index",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "ri:admin-line",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 300,
      "id": 302,
      "menuType": 0,
      "title": "menus.pureRole",
      "name": "SystemRole",
      "path": "/system/role/index",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "ri:admin-fill",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 300,
      "id": 303,
      "menuType": 0,
      "title": "menus.pureSystemMenu",
      "name": "SystemMenu",
      "path": "/system/menu/index",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "ep:menu",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 300,
      "id": 304,
      "menuType": 0,
      "title": "menus.pureDept",
      "name": "SystemDept",
      "path": "/system/dept/index",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "ri:git-branch-line",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 0,
      "id": 400,
      "menuType": 0,
      "title": "menus.pureSysMonitor",
      "name": "PureMonitor",
      "path": "/monitor",
      "component": "",
      "rank": 11,
      "redirect": "",
      "icon": "ep:monitor",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 400,
      "id": 401,
      "menuType": 0,
      "title": "menus.pureOnlineUser",
      "name": "OnlineUser",
      "path": "/monitor/online-user",
      "component": "monitor/online/index",
      "rank": null,
      "redirect": "",
      "icon": "ri:user-voice-line",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 400,
      "id": 402,
      "menuType": 0,
      "title": "menus.pureLoginLog",
      "name": "LoginLog",
      "path": "/monitor/login-logs",
      "component": "monitor/logs/login/index",
      "rank": null,
      "redirect": "",
      "icon": "ri:window-line",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 400,
      "id": 403,
      "menuType": 0,
      "title": "menus.pureOperationLog",
      "name": "OperationLog",
      "path": "/monitor/operation-logs",
      "component": "monitor/logs/operation/index",
      "rank": null,
      "redirect": "",
      "icon": "ri:history-fill",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 400,
      "id": 404,
      "menuType": 0,
      "title": "menus.pureSystemLog",
      "name": "SystemLog",
      "path": "/monitor/system-logs",
      "component": "monitor/logs/system/index",
      "rank": null,
      "redirect": "",
      "icon": "ri:file-search-line",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 0,
      "id": 500,
      "menuType": 0,
      "title": "menus.pureTabs",
      "name": "PureTabs",
      "path": "/tabs",
      "component": "",
      "rank": 12,
      "redirect": "",
      "icon": "ri:bookmark-2-line",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 500,
      "id": 501,
      "menuType": 0,
      "title": "menus.pureTabs",
      "name": "Tabs",
      "path": "/tabs/index",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": true,
      "showParent": false
    },
    {
      "parentId": 500,
      "id": 502,
      "menuType": 0,
      "title": "query传参模式",
      "name": "TabQueryDetail",
      "path": "/tabs/query-detail",
      "component": "",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "/tabs/index",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": false,
      "showParent": false
    },
    {
      "parentId": 500,
      "id": 503,
      "menuType": 0,
      "title": "params传参模式",
      "name": "TabParamsDetail",
      "path": "/tabs/params-detail/:id",
      "component": "params-detail",
      "rank": null,
      "redirect": "",
      "icon": "",
      "extraIcon": "",
      "enterTransition": "",
      "leaveTransition": "",
      "activePath": "/tabs/index",
      "auths": "",
      "frameSrc": "",
      "frameLoading": true,
      "keepAlive": false,
      "hiddenTag": false,
      "fixedTag": false,
      "showLink": false,
      "showParent": false
    }
  ]
}`

func Temp() {

	//ghub.Db.Model(&model.MenuBase{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")

	//err := ghub.Db.AutoMigrate(&model.Admin{})
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//err := ghub.Db.AutoMigrate(&model.MenuBase{}, &model.RolesBase{})
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//add()
	//add2()
}

func add() {

	// 解析JSON数据
	var adminMenu struct {
		Success bool        `json:"success"`
		Data    []MenuInput `json:"data"`
	}

	err := json.Unmarshal([]byte(jsonData), &adminMenu)
	if err != nil {
		log.Fatal(err)
	}
	var inputs = adminMenu.Data

	idMap := make(map[uint]uint)

	// 第一步：创建所有菜单并映射新旧ID
	for _, input := range inputs {
		newMenu := model.Menu{
			MenuType:        constants.MenuGenre(input.MenuType),
			Name:            input.Name,
			Path:            input.Path,
			Title:           input.Title,
			Component:       input.Component,
			Rank:            input.Rank,
			Redirect:        input.Redirect,
			Icon:            input.Icon,
			ExtraIcon:       input.ExtraIcon,
			EnterTransition: input.EnterTransition,
			LeaveTransition: input.LeaveTransition,
			ActivePath:      input.ActivePath,
			Auths:           input.Auths,
			FrameSrc:        input.FrameSrc,
			FrameLoading:    input.FrameLoading,
			KeepAlive:       input.KeepAlive,
			HiddenTag:       input.HiddenTag,
			FixedTag:        input.FixedTag,
			ShowLink:        input.ShowLink,
			ShowParent:      input.ShowParent,
		}
		result := ghub.Db.Create(&newMenu)
		if result.Error != nil {
			log.Fatalf("Failed to create menu: %v", result.Error)
		}
		idMap[input.OriginalID] = newMenu.ID
	}

	// 第二步：更新所有菜单的 ParentID
	for _, input := range inputs {
		id := idMap[input.OriginalID]
		if id != 0 && &id != nil {
			originalParentId := idMap[input.ParentID]
			if originalParentId != 0 && &originalParentId != nil {
				ghub.Db.Model(&model.Menu{}).Where("id = ?", id).Update("parent_id", &originalParentId)
			}
		}

	}
}

var jsonData2 = `{
  "success": true,
  "data": [
    {
      "name": "杭州总公司",
      "parentId": 0,
      "id": 100,
      "sort": 0,
      "phone": "15888888888",
      "principal": "呈轩",
      "email": "gs2m56_om0@yahoo.cn",
      "status": 1,
      "type": 1,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "郑州分公司",
      "parentId": 100,
      "id": 101,
      "sort": 1,
      "phone": "15888888888",
      "principal": "浩轩",
      "email": "p1pscp_pp8@gmail.com",
      "status": 1,
      "type": 2,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "研发部门",
      "parentId": 101,
      "id": 103,
      "sort": 1,
      "phone": "15888888888",
      "principal": "睿渊",
      "email": "nlrlqi_nev@126.com",
      "status": 1,
      "type": 3,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "市场部门",
      "parentId": 102,
      "id": 108,
      "sort": 1,
      "phone": "15888888888",
      "principal": "懿轩",
      "email": "jlbi4665@21cn.com",
      "status": 1,
      "type": 3,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "深圳分公司",
      "parentId": 100,
      "id": 102,
      "sort": 2,
      "phone": "15888888888",
      "principal": "健雄",
      "email": "my272@gmail.com",
      "status": 1,
      "type": 2,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "市场部门",
      "parentId": 101,
      "id": 104,
      "sort": 2,
      "phone": "15888888888",
      "principal": "文韬",
      "email": "uclgm3.scd@21cn.com",
      "status": 1,
      "type": 3,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "财务部门",
      "parentId": 102,
      "id": 109,
      "sort": 2,
      "phone": "15888888888",
      "principal": "涛",
      "email": "fextfv.g3k@tom.com",
      "status": 1,
      "type": 3,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "测试部门",
      "parentId": 101,
      "id": 105,
      "sort": 3,
      "phone": "15888888888",
      "principal": "正豪",
      "email": "k1zu0c_i8w79@vip.qq.com",
      "status": 0,
      "type": 3,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "财务部门",
      "parentId": 101,
      "id": 106,
      "sort": 4,
      "phone": "15888888888",
      "principal": "睿渊",
      "email": "i1csbt81@tom.com",
      "status": 1,
      "type": 3,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    },
    {
      "name": "运维部门",
      "parentId": 101,
      "id": 107,
      "sort": 5,
      "phone": "15888888888",
      "principal": "凯瑞",
      "email": "ggqk1z_fnq21@gmail.com",
      "status": 0,
      "type": 3,
      "createTime": 1605456000000,
      "remark": "这里是备注信息这里是备注信息这里是备注信息这里是备注信息"
    }
  ]
}`

type DeptInput struct {
	OriginalID uint   `json:"id"`
	ParentID   uint   `json:"parentId"`
	Name       string `json:"name"`
	Principal  string `json:"principal"`
	Email      string `json:"email"`
	Mobile     string `json:"phone;"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	Remark     string `json:"remark"`
}

func add2() {
	// 解析JSON数据
	var adminMenu struct {
		Success bool        `json:"success"`
		Data    []DeptInput `json:"data"`
	}

	err := json.Unmarshal([]byte(jsonData2), &adminMenu)
	if err != nil {
		log.Fatal(err)
	}
	var inputs = adminMenu.Data

	idMap := make(map[uint]uint)

	// 第一步：创建所有菜单并映射新旧ID
	for _, input := range inputs {
		newMenu := model.Dept{
			Name:      input.Name,
			Principal: input.Principal,
			Email:     input.Email,
			Mobile:    input.Mobile,
			Sort:      input.Sort,
			Remark:    input.Remark,
			Status:    input.Status == 1,
		}
		result := ghub.Db.Create(&newMenu)
		if result.Error != nil {
			log.Fatalf("Failed to create menu: %v", result.Error)
		}
		idMap[input.OriginalID] = newMenu.ID
	}

	// 第二步：更新所有菜单的 ParentID
	for _, input := range inputs {
		id := idMap[input.OriginalID]
		if id != 0 && &id != nil {
			originalParentId := idMap[input.ParentID]
			if originalParentId != 0 && &originalParentId != nil {
				ghub.Db.Model(&model.Dept{}).Where("id = ?", id).Update("parent_id", &originalParentId)
			}
		}

	}
}
