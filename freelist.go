package main

func newFreeList(maxpage pgnum) *freelist {
	return &freelist{
		maxPage: maxpage,
		releasedPages: []pgnum{},
	}
}

func (fr *freelist) getNextPage() pgnum {
	if len(fr.releasedPages) != 0 {
		pageId := fr.releasedPages[len(fr.releasedPages)-1]
		fr.releasedPages = fr.releasedPages[:len(fr.releasedPages)-1]
		return pageId
	}
	fr.maxPage += 1
	return fr.maxPage
}

func (fr *freelist) releasePage(page pgnum) {
	fr.releasedPages = append(fr.releasedPages, page)
}