package preloads

type getKey[Key comparable, T any] func(t T) Key

type PreloadArg[Target, PreloadItem any, Key comparable] struct {
	KeyByItem   getKey[Key, PreloadItem]
	KeyByTarget getKey[Key, Target]
	SetItem     func(target *Target, items PreloadItem)
}

// Preload is a function to preload one item to many targets
func Preload[Target any, PreloadItem any, Key comparable](
	targets []Target,
	preloadItems []PreloadItem,
	arg PreloadArg[Target, PreloadItem, Key],
) []Target {
	mapper := make(map[Key]PreloadItem, len(preloadItems))
	for _, item := range preloadItems {
		mapper[arg.KeyByItem(item)] = item
	}

	for i := range targets {
		key := arg.KeyByItem(preloadItems[i])
		arg.SetItem(&targets[i], mapper[key])
	}

	return targets
}

type PreloadManyArg[Target, PreloadItem any, Key comparable] struct {
	KeyByItem   getKey[Key, PreloadItem]
	KeyByTarget getKey[Key, Target]
	SetItem     func(target *Target, items []PreloadItem)
}

// PreloadsMany is a function to preload many items to many targets
func PreloadsMany[Target any, PreloadItem any, Key comparable](
	targets []Target,
	preloadItems []PreloadItem,
	arg PreloadManyArg[Target, PreloadItem, Key],
) []Target {
	mapper := make(map[Key][]PreloadItem, len(preloadItems))
	for _, item := range preloadItems {
		key := arg.KeyByItem(item)
		mapper[key] = append(mapper[key], item)
	}

	for i := range targets {
		key := arg.KeyByTarget(targets[i])
		arg.SetItem(&targets[i], mapper[key])
	}

	return targets
}
