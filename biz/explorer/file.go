package explorer

import "net/http"

// Download 签名的匿名文件下载
func (service *FileAnonymousGetService) Download(ctx context.Context, c *gin.Context) serializer.Response {
	fs, err := filesystem.NewAnonymousFileSystem()
	if err != nil {
		return serializer.Err(serializer.CodeCreateFSError, "", err)
	}
	defer fs.Recycle()

	// 查找文件
	err = fs.SetTargetFileByIDs([]uint{service.ID})
	if err != nil {
		return serializer.Err(serializer.CodeNotSet, err.Error(), err)
	}

	// 获取文件流
	rs, err := fs.GetDownloadContent(ctx, 0)
	defer rs.Close()
	if err != nil {
		return serializer.Err(serializer.CodeNotSet, err.Error(), err)
	}

	// 发送文件
	http.ServeContent(c.Writer, c.Request, service.Name, fs.FileTarget[0].UpdatedAt, rs)

	return serializer.Response{
		Code: 0,
	}
}

// DownloadArchived 通过预签名 URL 打包下载
func (service *ArchiveService) DownloadArchived(ctx context.Context, c *gin.Context) serializer.Response {
	userRaw, exist := cache.Get("archive_user_" + service.ID)
	if !exist {
		return serializer.Err(serializer.CodeNotFound, "Archive session not exist", nil)
	}
	user := userRaw.(model.User)

	// 创建文件系统
	fs, err := filesystem.NewFileSystem(&user)
	if err != nil {
		return serializer.Err(serializer.CodeCreateFSError, "", err)
	}
	defer fs.Recycle()

	// 查找打包的临时文件
	archiveSession, exist := cache.Get("archive_" + service.ID)
	if !exist {
		return serializer.Err(serializer.CodeNotFound, "Archive session not exist", nil)
	}

	// 开始打包
	c.Header("Content-Disposition", "attachment;")
	c.Header("Content-Type", "application/zip")
	itemService := archiveSession.(ItemIDService)
	items := itemService.Raw()
	ctx = context.WithValue(ctx, fsctx.GinCtx, c)
	err = fs.Compress(ctx, c.Writer, items.Dirs, items.Items, true)
	if err != nil {
		return serializer.Err(serializer.CodeNotSet, "Failed to compress file", err)
	}

	return serializer.Response{
		Code: 0,
	}
}
