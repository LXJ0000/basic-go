package handler

import (
	"github.com/gin-gonic/gin"
	"sort"
	g "webook-server/internal/global"
	"webook-server/internal/middleware"
	"webook-server/internal/model"
	"webook-server/internal/repository"
)

type MenuHandler struct {
	repo repository.MenuRepository
}

func NewMenuHandler(repo repository.MenuRepository) MenuHandler {
	return MenuHandler{repo: repo}
}

func (h *MenuHandler) InitRouter(r *gin.Engine) {
	base := r.Group("/api/menu").Use(middleware.JwtAuthMiddleware())
	base.GET("/list", h.GetTreeList)
	base.GET("/user/list", h.GetUserMenu)

}

func (h *MenuHandler) GetTreeList(c *gin.Context) {

}

func (h *MenuHandler) GetUserMenu(c *gin.Context) {
	type Resp struct {
		menu []MenuTreeVO
	}
	menus, err := h.repo.GetAllMenuList(c)
	if err != nil {
		ReturnFail(c, g.ErrDbOp, err.Error())
	}
	ReturnSuccess(c, menus2MenuVos(menus))
}

type MenuTreeVO struct {
	model.Menu
	Children []MenuTreeVO `json:"children"`
}

func menus2MenuVos(menus []model.Menu) []MenuTreeVO {
	res := make([]MenuTreeVO, 0)
	firstLevelMenu := getFirstLevelMenu(menus)
	childrenMap := getMenuChildrenMap(menus)
	for _, first := range firstLevelMenu {
		menu := MenuTreeVO{Menu: first}
		for _, child := range childrenMap[first.Id] {
			menu.Children = append(menu.Children, MenuTreeVO{Menu: child})
		}
		delete(childrenMap, first.Id)
		res = append(res, menu)
	}
	sortMenu(res)
	return res
}

func getFirstLevelMenu(menus []model.Menu) []model.Menu {
	firstLevelMenu := make([]model.Menu, 0)
	for _, menu := range menus {
		if menu.ParentId == 0 {
			firstLevelMenu = append(firstLevelMenu, menu)
		}
	}
	return firstLevelMenu
}

func getMenuChildrenMap(menus []model.Menu) map[int64][]model.Menu {
	childrenMap := make(map[int64][]model.Menu)
	for _, menu := range menus {
		if menu.ParentId != 0 {
			childrenMap[menu.ParentId] = append(childrenMap[menu.ParentId], menu)
		}
	}
	return childrenMap
}

func sortMenu(menus []MenuTreeVO) {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].OrderNum < menus[j].OrderNum
	})
	for i := range menus {
		sort.Slice(menus[i].Children, func(j, k int) bool {
			return menus[i].Children[j].OrderNum < menus[i].Children[k].OrderNum
		})
	}
}
