package system

import (
	"github.com/562589540/jono-gin/ghub/glibrary/gdto"
	"github.com/562589540/jono-gin/ghub/glibrary/gres"
	"github.com/562589540/jono-gin/internal/app/system/dto"
	"github.com/562589540/jono-gin/internal/app/system/service"
	"github.com/gin-gonic/gin"
)

type JobApi struct {
	jobService service.IJobService
}

func NewJobApi(jobService service.IJobService) *JobApi {
	return &JobApi{
		jobService: jobService,
	}
}

func (m JobApi) Create(c *gin.Context, req dto.JobAddReq) (any, error) {
	if err := m.jobService.Create(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "创建成功",
	}, nil
}

func (m JobApi) List(c *gin.Context, req dto.JobSearchReq) (any, error) {
	list, total, err := m.jobService.List(c, &req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "查询成功",
		Data: gdto.ListRes{
			List:  list,
			Total: total,
		},
	}, nil
}

func (m JobApi) JobLog(c *gin.Context, req dto.JobLogReq) (any, error) {
	list, total, err := m.jobService.JobLog(c, &req)
	if err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "查询成功",
		Data: gdto.ListRes{
			List:  list,
			Total: total,
		},
	}, nil
}

func (m JobApi) Update(c *gin.Context, req dto.JobUpdateReq) (any, error) {
	if err := m.jobService.Update(c, &req); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "更新成功",
	}, nil
}

func (m JobApi) Delete(c *gin.Context, req gdto.IDSReq) (any, error) {
	if err := m.jobService.Delete(c, req.IDS); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}
func (m JobApi) DeleteJobLog(c *gin.Context, req gdto.IDSReq) (any, error) {
	if err := m.jobService.DeleteJobLog(c, req.IDS); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}

func (m JobApi) DeleteJobLogAll(c *gin.Context, req gdto.IDReq) (any, error) {
	if err := m.jobService.DeleteJobLogAll(c, req.ID); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "删除成功",
	}, nil
}

func (m JobApi) Once(c *gin.Context, req gdto.IDReq) (any, error) {
	if err := m.jobService.Once(c, req.ID); err != nil {
		return nil, err
	}
	return gres.Response{
		Message: "执行成功",
	}, nil
}
