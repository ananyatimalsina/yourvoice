package utils

// Size options
const (
	SizeSmall  = "small"
	SizeMedium = "medium"
	SizeLarge  = "large"
	SizeFull   = "full"
)

// Input types
const (
	InputTypeText     = "text"
	InputTypeSelect   = "select"
	InputTypeMultiple = "multiple"
	InputTypeDate     = "date"
)

// Dialog types
const (
	DialogTypeConfirm     = "confirm"
	DialogTypeAcknowledge = "acknowledge"
)

// Page size options
const (
	PageSizeCompact     = 25
	PageSizeComfortable = 100
	PageSizeExpanded    = 500
)

// Action style options
const (
	ActionStylePrimary   = "bg-blue-600 text-white hover:bg-blue-500"
	ActionStyleSecondary = "bg-slate-600 text-slate-200 hover:bg-slate-500"
	ActionStyleGhost     = "text-blue-400 hover:text-blue-300 hover:bg-slate-700"
	ActionStyleWarning   = "bg-yellow-500 hover:bg-yellow-600 text-white"
	ActionStyleDanger    = "bg-red-600 hover:bg-red-700 text-white"
	ActionStyleGood      = "bg-green-600 hover:bg-green-500 text-white"

	// Additional color variants for dialogs
	ActionStylePurple = "bg-purple-600 hover:bg-purple-500 text-white"
	ActionStyleIndigo = "bg-indigo-600 hover:bg-indigo-500 text-white"
	ActionStylePink   = "bg-pink-600 hover:bg-pink-500 text-white"
	ActionStyleOrange = "bg-orange-600 hover:bg-orange-500 text-white"
	ActionStyleTeal   = "bg-teal-600 hover:bg-teal-500 text-white"

	// Muted variants
	ActionStyleMuted = "bg-slate-500 text-slate-100 hover:bg-slate-400"
	ActionStyleLight = "bg-slate-200 text-slate-800 hover:bg-slate-100"
	ActionStyleDark  = "bg-slate-900 text-slate-100 hover:bg-slate-800"

	// Outline variants
	ActionStyleOutlinePrimary   = "border-2 border-blue-600 text-blue-400 hover:bg-blue-600 hover:text-white"
	ActionStyleOutlineSecondary = "border-2 border-slate-600 text-slate-300 hover:bg-slate-600 hover:text-white"
	ActionStyleOutlineDanger    = "border-2 border-red-600 text-red-400 hover:bg-red-600 hover:text-white"
	ActionStyleOutlineGood      = "border-2 border-green-600 text-green-400 hover:bg-green-600 hover:text-white"
)

// Indicator variant options
const (
	IndicatorStyleSuccess = "bg-green-900/50 text-green-300 border border-green-700"
	IndicatorStyleWarning = "bg-yellow-900/50 text-yellow-300 border border-yellow-700"
	IndicatorStyleError   = "bg-red-900/50 text-red-300 border border-red-700"
	IndicatorStyleInfo    = "bg-blue-900/50 text-blue-300 border border-blue-700"
	IndicatorStyleDefault = "default"
)
