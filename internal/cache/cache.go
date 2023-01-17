package cache

func MustInit(interval int) {
	if err := runAppConfig(interval); err != nil {
		panic("run appConfig cache err: " + err.Error())
	}
}

func runAppConfig(interval int) error {
	//cache 使用示例
	//ticker := time.NewTicker(time.Duration(interval) * time.Second)
	//ctx, _ := context.WithCancel(context.Background()) //nolint
	//syncData(ctx)
	//go func() {
	//	for {
	//		select {
	//		case <-ticker.C:
	//			syncData(ctx)
	//		case <-ctx.Done():
	//			glog.Warnf(ctx, "app config sync shutdown ")
	//			ticker.Stop()
	//			return
	//		}
	//	}
	//}()
	return nil
}
