package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/mlogclub/mlog/services/cache"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/mlog/controllers/render"
	"github.com/mlogclub/mlog/model"
	"github.com/mlogclub/mlog/services"
	"github.com/mlogclub/mlog/utils"
)

type IndexController struct {
	Ctx            iris.Context
	TagService     *services.TagService
	ArticleService *services.ArticleService
	TopicService   *services.TopicService
}

func (this *IndexController) Any() mvc.View {
	categories := cache.CategoryCache.GetAllCategories()
	activeUsers := cache.UserCache.GetActiveUsers()
	activeTags := cache.TagCache.GetActiveTags()

	articles, _ := this.ArticleService.QueryCnd(simple.NewQueryCnd("status = ?", model.ArticleStatusPublished).
		Order("id desc").Size(20))

	topics, _ := this.TopicService.QueryCnd(simple.NewQueryCnd("status = ?", model.TopicStatusOk).Order("id desc").Size(10))

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			utils.GlobalFieldSiteDescription: "M-LOG分享",
			utils.GlobalFieldSiteKeywords:    "Go中国，Golang, Golang学习,Golang教程,Golang社区,Go基金会,Go语言中文网,Go,Go语言,主题,资源,文章,图书,开源项目,M-LOG",
			"Categories":                     categories,
			"Articles":                       render.BuildArticles(articles),
			"Topics":                         render.BuildTopics(topics),
			"ActiveUsers":                    render.BuildUsers(activeUsers),
			"ActiveTags":                     render.BuildTags(activeTags),
		},
	}
}

func (this *IndexController) GetAbout() mvc.View {
	return mvc.View{
		Name: "about.html",
	}
}
