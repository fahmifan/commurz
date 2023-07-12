package preloads

type getRef[Ref comparable, T any] func(t T) Ref

type Preload[Target, PreloadItem any, RefKey comparable] struct {
	Targets    []Target
	RefItem    getRef[RefKey, PreloadItem]
	RefTarget  getRef[RefKey, Target]
	SetItem    func(target *Target, items PreloadItem)
	FetchItems func() ([]PreloadItem, error)
}

func (arg Preload[Target, PreloadItem, RefKey]) Preload() ([]Target, error) {
	preloadItems, err := arg.FetchItems()
	if err != nil {
		return nil, err
	}

	mapper := make(map[RefKey]PreloadItem, len(preloadItems))
	for _, item := range preloadItems {
		mapper[arg.RefItem(item)] = item
	}

	for i := range arg.Targets {
		key := arg.RefTarget(arg.Targets[i])
		arg.SetItem(&arg.Targets[i], mapper[key])
	}

	return arg.Targets, nil
}

type PreloadMany[Target, PreloadItem any, RefKey comparable] struct {
	Targets    []Target
	RefItem    getRef[RefKey, PreloadItem]
	RefTarget  getRef[RefKey, Target]
	SetItem    func(target *Target, items []PreloadItem)
	FetchItems func() ([]PreloadItem, error)
}

func (arg PreloadMany[Target, PreloadItem, RefKey]) Preload() ([]Target, error) {
	preloadItems, err := arg.FetchItems()
	if err != nil {
		return nil, err
	}

	mapper := make(map[RefKey][]PreloadItem, len(preloadItems))
	for _, item := range preloadItems {
		key := arg.RefItem(item)
		mapper[key] = append(mapper[key], item)
	}

	for i := range arg.Targets {
		key := arg.RefTarget(arg.Targets[i])
		arg.SetItem(&arg.Targets[i], mapper[key])
	}

	return arg.Targets, nil
}
