package service

import (
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strconv"
	"strings"
	"yanhaiproject/core"
	"yanhaiproject/model"
	"yanhaiproject/tool"
	log "github.com/sirupsen/logrus"
)




type PictureService struct {
}

/**
 * @author ZhangYiBo
 * @description  将图片id转换成url
 * @date 10:54 上午 2021/9/16
 * @param picId 传入的图片id
 * @return 图片URL
 **/
func (PictureService) PicIdToURL(picId int) string {
	var picture model.Picture
	core.DB.First(&picture, picId)
	return core.URL_PREFIX + picture.PicURL
}


/**
 * @author ZhangYiBo
 * @description  将以|格式的id转换成URL
 * @date 10:55 上午 2021/9/16
 * @param picIds 以|隔开的字符串
 * @return URL的切片
 **/
func (PictureService) PicIdsToURL(picIds string) []string  {
	//TODO 返回字符串切片类型，带测试
	allPicIds := strings.Split(picIds, "|")
	iPicIds := make([]int, len(allPicIds))
	for index, str := range allPicIds{
		iPicIds[index],_ = strconv.Atoi(str)
	}
	var topicPics []model.Picture
	core.DB.Find(&topicPics, iPicIds)
	retPictureURLs := make([]string, len(topicPics))
	for index, pic := range topicPics{
		retPictureURLs[index] = core.URL_PREFIX + pic.PicURL
	}
	return retPictureURLs
}


/**
 * @author ZhangYiBo
 * @description  存储上传的图片
 * @date 9:57 上午 2021/9/17
 * @param context 上下文
 * @return mess 回传消息 isSuccess是否上传成功
 **/
func StorePictures(context *gin.Context) (mess string, isSuccess bool, ids []int, urls []string) {
	//接口测试通过
	//提前构造存储目录
	runenv := core.ApplicationConfig.GetString("runenv")
	var dirPath string
	if runenv == "dev" {
		dirPath = core.ApplicationConfig.GetString("picpath.devpath")
	} else if runenv == "prod" {
		dirPath = core.ApplicationConfig.GetString("picpath.prodpath")
	}else {
		log.Debug("初始化文件服务器路径出错")
		mess = "文件服务器出错,获取配置信息出错"
		isSuccess = false
		return
	}

	form, _ := context.MultipartForm()
	files := form.File["files"]
	//创建等同长度的数据用来存储文件存储成功之后返回的id
	picIds := make([]int, len(files))
	urls = make([]string,len(files))
	for index, file := range files{
		//获取判断扩展名是否合法
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg": true,
			".png": true,
			".gif": true,
			".jpeg": true,
		}

		if _, isAllow := allowExtMap[extName]; !isAllow {
			log.Debug("文件后缀名不合法")
			mess = "文件名称不合法"
			isSuccess = false
			ids = nil
			urls = nil
			return
		}

		//构造存储目录
		day := tool.GetDay()
		createDirPath := dirPath + day

		if err:= os.MkdirAll(createDirPath, 0777); err != nil{
			log.Error(err)
			mess = "文件服务器出错，创建文件夹出错"
			isSuccess = false
			ids = nil
			urls = nil
			return
		}

		//重新生成文件名 运行速度快 需要获取纳秒时间才行
		fileUnixName := strconv.FormatInt(tool.GetUnixNano(), 10)
		saveDst := path.Join(createDirPath, fileUnixName + extName)
		err := context.SaveUploadedFile(file, saveDst)
		//记录数据库文件存储
		if err != nil {
			//出错处理
			mess = "存储失败"
			isSuccess = false
			ids = nil
			urls = nil
			log.Error(err)
			return
		}else {
			//保存图片信息至数据库
			var picture model.Picture
			picture.PicURL = path.Join(day, fileUnixName + extName)
			log.Info("保存在数据库中的图片地址将为")
			log.Info(picture.PicURL)
			core.DB.Create(&picture)
			picIds[index] = picture.PicId
			urls[index] = core.URL_PREFIX + picture.PicURL
		}
	}
	return "成功",true, picIds, urls
}
