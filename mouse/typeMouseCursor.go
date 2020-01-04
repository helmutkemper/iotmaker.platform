package mouse

type CursorType int

func (el CursorType) String() string {
	return cursorTypesWebBrowser[el]
}

var cursorTypesWebBrowser = [...]string{
	"default",
	"arrowBottomLeft",
	"arrowBottomRight",
	"arrowDown",
	"arrowCollapse",
	"arrowCollapseAll",
	"arrowCollapseDown",
	"arrowCollapseHorizontal",
	"arrowCollapseLeft",
	"arrowCollapseRight",
	"arrowCollapseUp",
	"arrowCollapseVertical",
	"arrowExpand",
	"arrowExpandAll",
	"arrowExpandDown",
	"arrowExpandHorizontal",
	"arrowExpandLeft",
	"arrowExpandRight",
	"arrowExpandUp",
	"arrowExpandVertical",
	"arrowHorizontalLock",
	"arrowLeft",
	"arrowLeftRight",
	"arrowRight",
	"arrowTopLeft",
	"arrowTopLeftBottomRight",
	"arrowTopRight",
	"arrowTopRightBottomLeft",
	"arrowUp",
	"arrowUpDown",
}

const (
	KCursorDefault CursorType = iota
	KCursorArrowBottomLeft
	KCursorArrowBottomRight
	KCursorArrowDown
	KCursorArrowCollapse
	KCursorArrowCollapseAll
	KCursorArrowCollapseDown
	KCursorArrowCollapseHorizontal
	KCursorArrowCollapseLeft
	KCursorArrowCollapseRight
	KCursorArrowCollapseUp
	KCursorArrowCollapseVertical
	KCursorArrowExpand
	KCursorArrowExpandAll
	KCursorArrowExpandDown
	KCursorArrowExpandHorizontal
	KCursorArrowExpandLeft
	KCursorArrowExpandRight
	KCursorArrowExpandUp
	KCursorArrowExpandVertical
	KCursorArrowHorizontalLock
	KCursorArrowLeft
	KCursorArrowLeftRight
	KCursorArrowRight
	KCursorArrowTopLeft
	KCursorArrowTopLeftBottomRight
	KCursorArrowTopRight
	KCursorArrowTopRightBottomLeft
	KCursorArrowUp
	KCursorArrowUpDown
)
