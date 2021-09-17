package service

import (
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strconv"
	"yanhaiproject/core"
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
	//FIXME 接口待完成
	return "http://www.xxx.xxx"
}


/**
 * @author ZhangYiBo
 * @description  将以|格式的id转换成URL
 * @date 10:55 上午 2021/9/16
 * @param picIds 以|隔开的字符串
 * @return URL的切片
 **/
func (PictureService) PicIdsToURL(picIds string) []string  {
	//FIXME 检查返回值为切片类型，尚不确定切片类型的返回值该怎么返回
	return []string{}
}


/**
 * @author ZhangYiBo
 * @description  存储上传的图片
 * @date 9:57 上午 2021/9/17
 * @param context 上下文
 * @return mess 回传消息 isSuccess是否上传成功
 **/
func StorePictures(context *gin.Context) (mess string, isSuccess bool) {
	//FIXME 等待接口测试
	//提前构造存储目录
	runenv := core.ApplicationConfig.GetString("runenv")
	var dirPath string
	if runenv == "dev" {
		dirPath = core.ApplicationConfig.GetString("devpath")
	} else if dirPath == "prod" {
		dirPath = core.ApplicationConfig.GetString("prodpath")
	}else {
		mess = "文件服务器出错"
		isSuccess = false
		return
	}

	form, _ := context.MultipartForm()
	files := form.File["files"]
	for _, file := range files{
		//获取判断扩展名是否合法
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg": true,
			".png": true,
			".gif": true,
			".jpeg": true,
		}

		if _, isAllow := allowExtMap[extName]; !isAllow {
			mess = "文件名称不合法"
			isSuccess = false
			return
		}

		//构造存储目录
		day := tool.GetDay()
		dirPath += day

		if err:= os.MkdirAll(dirPath, 0666); err != nil{
			log.Error(err)
			mess = "文件服务器出错"
			isSuccess = false
			return
		}

		//重新生成文件名
		fileUnixName := strconv.FormatInt(tool.GetUnix(), 10)
		saveDst := path.Join(dirPath, fileUnixName + extName)
		err := context.SaveUploadedFile(file, saveDst)
		//FIXME 记录数据库文件存储
		if err != nil {
			//出错处理
			mess = "存储失败"
			isSuccess = false
			log.Error(err)
			return
		}else {
			//保存图片信息至数据库
		}
	}
	return "成功",true
}
