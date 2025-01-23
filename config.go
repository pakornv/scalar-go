package scalar

// Config represents the configuration options for the reference
// For more information, see: https://github.com/scalar/scalar/blob/main/documentation/configuration.md
type Config struct {
	// The Swagger/OpenAPI spec to render
	// Spec Spec `json:"spec"`

	// A string to use one of the color presets
	Theme Theme `json:"theme,omitempty"`

	// The layout to use for the references
	Layout Layout `json:"layout,omitempty"`

	// URL to a request proxy for the API client
	ProxyURL string `json:"proxyUrl,omitempty"`

	// Whether the spec input should show
	IsEditable *bool `json:"isEditable,omitempty"`

	// Whether to show the sidebar
	ShowSidebar *bool `json:"showSidebar,omitempty"`

	// Whether to show models in the sidebar, search, and content.
	// Default: false
	HideModels *bool `json:"hideModels,omitempty"`

	// Whether to show the "Download OpenAPI Document" button
	// Default: false
	HideDownloadButton *bool `json:"hideDownloadButton,omitempty"`

	// Whether to show the "Test Request" button
	// Default: false
	HideTestRequestButton *bool `json:"hideTestRequestButton,omitempty"`

	// Whether to show the sidebar search bar
	// Default: false
	HideSearch *bool `json:"hideSearch,omitempty"`

	// Whether dark mode is on or off initially (light mode)
	DarkMode *bool `json:"darkMode,omitempty"`

	// forceDarkModeState makes it always this state no matter what
	ForceDarkModeState ForceDarkModeState `json:"forceDarkModeState,omitempty"`

	// Whether to show the dark mode toggle
	HideDarkModeToggle *bool `json:"hideDarkModeToggle,omitempty"`

	// Key used with CTRL/CMD to open the search modal (defaults to 'k' e.g. CMD+k)
	SearchHotKey string `json:"searchHotKey,omitempty"` // one character from a-z

	// if used, passed data will be added to the HTML header
	// See: https://unhead.unjs.io/usage/composables/use-seo-meta
	MetaData map[string]string `json:"metaData,omitempty"`

	// Path to a favicon image
	// Default: undefined
	// Example: '/favicon.svg'
	Favicon string `json:"favicon,omitempty"`

	// List of httpsnippet clients to hide from the clients menu
	// By default hides Unirest, pass empty slice to show all clients
	HiddenClients []HTTPClient `json:"hiddenClients,omitempty"`

	// Determine the HTTP client that's selected by default
	DefaultHTTPClient *HTTPClientState `json:"defaultHttpClient,omitempty"`

	// Custom CSS to be added to the page
	CustomCSS string `json:"customCss,omitempty"`

	// OnSpecUpdate is fired on spec/swagger content change
	// OnSpecUpdate any `json:"onSpecUpdate,omitempty"`

	// Prefill authentication
	Authentication map[string]any `json:"authentication,omitempty"`

	// Route using paths instead of hashes, your server MUST support this
	// for example vue router needs a catch all so any subpaths are included
	//
	// Example:
	// '/standalone-api-reference/:custom(.*)?'
	//
	// Experimental
	// Default: undefined
	PathRouting *PathRouting `json:"pathRouting,omitempty"`

	// Custom slug generation functions
	// GenerateHeadingSlug   any `json:"generateHeadingSlug,omitempty"`
	// GenerateModelSlug     any `json:"generateModelSlug,omitempty"`
	// GenerateTagSlug       any `json:"generateTagSlug,omitempty"`
	// GenerateOperationSlug any `json:"generateOperationSlug,omitempty"`
	// GenerateWebhookSlug   any `json:"generateWebhookSlug,omitempty"`

	// The baseServerURL is used when the spec servers are relative paths and we are using SSR
	// Default: undefined
	// Example: 'http://localhost:3000'
	BaseServerURL string `json:"baseServerURL,omitempty"`

	// List of servers to override the openapi spec servers
	Server []Server `json:"servers,omitempty"`

	// We're using Inter and JetBrains Mono as the default fonts. If you want to use your own fonts, set this to false.
	// Default: true
	WithDefaultFonts *bool `json:"withDefaultFonts,omitempty"`

	// By default we only open the relevant tag based on the url, however if you want all the tags open by default then set this configuration option
	// Default: false
	DefaultOpenAllTags *bool `json:"defaultOpenAllTags,omitempty"`

	// Sort tags alphabetically or with a custom sort function
	// TagsSorter any `json:"tagsSorter,omitempty"`

	// Sort operations alphabetically, by method or with a custom sort function
	// OperationsSorter any `json:"operationsSorter,omitempty"`

	// Whether to show the client button from the reference sidebar and modal
	// Default: false
	HideClientButton *bool `json:"hideClientButton,omitempty"`

	// Additional properties not in the official configuration
	Title string `json:"-"` // Title of the page
	CDN   string `json:"-"` // URL to the Scalar CDN
}

// Theme represents available themes
type Theme string

const (
	ThemeAlternate  Theme = "alternate"
	ThemeDefault    Theme = "default"
	ThemeMoon       Theme = "moon"
	ThemePurple     Theme = "purple"
	ThemeSolarized  Theme = "solarized"
	ThemeBluePlanet Theme = "bluePlanet"
	ThemeDeepSpace  Theme = "deepSpace"
	ThemeSaturn     Theme = "saturn"
	ThemeKepler     Theme = "kepler"
	ThemeElysiajs   Theme = "elysiajs"
	ThemeFastify    Theme = "fastify"
	ThemeMars       Theme = "mars"
	ThemeNone       Theme = "none"
)

// Layout represents available layouts
type Layout string

const LayoutClassic Layout = "classic"
const LayoutModern Layout = "modern"

type ForceDarkModeState string

const (
	Dark  ForceDarkModeState = "dark"
	Light ForceDarkModeState = "light"
)

// HTTPClient represents available clients
type HTTPClient string

const (
	CLibcurl             HTTPClient = "c/libcurl"
	ClojureHTTP          HTTPClient = "clojure/clj_http"
	CSharpHTTPClient     HTTPClient = "csharp/httpclient"
	CSharpRestSharp      HTTPClient = "csharp/restsharp"
	GoNative             HTTPClient = "go/native"
	HTTPOnePOne          HTTPClient = "http/http1.1"
	JavaAsyncHTTP        HTTPClient = "java/asynchttp"
	JavaNetHTTP          HTTPClient = "java/nethttp"
	JavaOkHTTP           HTTPClient = "java/okhttp"
	JavaUnirest          HTTPClient = "java/unirest"
	JSAxios              HTTPClient = "js/axios"
	JSFetch              HTTPClient = "js/fetch"
	JSJQuery             HTTPClient = "js/jquery"
	JSOfetch             HTTPClient = "js/ofetch"
	JSXHR                HTTPClient = "js/xhr"
	KotlinOkHTTP         HTTPClient = "kotlin/okhttp"
	NodeAxios            HTTPClient = "node/axios"
	NodeFetch            HTTPClient = "node/fetch"
	NodeOfetch           HTTPClient = "node/ofetch"
	NodeUndici           HTTPClient = "node/undici"
	ObjcNSURLSession     HTTPClient = "objc/nsurlsession"
	OCamlCoHTTP          HTTPClient = "ocaml/cohttp"
	PHPCurl              HTTPClient = "php/curl"
	PHPGuzzle            HTTPClient = "php/guzzle"
	PowerShellRestmethod HTTPClient = "powershell/restmethod"
	PowerShellWebRequest HTTPClient = "powershell/webrequest"
	PythonPython3        HTTPClient = "python/python3"
	PythonRequests       HTTPClient = "python/requests"
	RHttr                HTTPClient = "r/httr"
	RubyNative           HTTPClient = "ruby/native"
	ShellCurl            HTTPClient = "shell/curl"
	ShellHTTPie          HTTPClient = "shell/httpie"
	ShellWGet            HTTPClient = "shell/wget"
	SwiftNSURLSession    HTTPClient = "swift/nsurlsession"
	DartHTTP             HTTPClient = "dart/http"
)

type HTTPClientState struct {
	TargetKey string `json:"targetKey"`
	ClientKey string `json:"clientKey"`
}

type PathRouting struct {
	BasePath string `json:"basePath"`
}

type Server struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}
