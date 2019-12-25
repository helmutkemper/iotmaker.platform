package svg

type A struct {
	// en: Instructs browsers to download a URL instead of navigating to it, so the
	// user will be prompted to save it as a local file.
	//     Value type: <string> ; Default value: none; Animatable: no
	Download string

	// en: The URL or URL fragment the hyperlink points to.
	//     Value type: <URL> ; Default value: none; Animatable: yes
	HRef string

	// en: The human language of the URL or URL fragment that the hyperlink points to.
	//     Value type: <string> ; Default value: none; Animatable: yes
	HRefLang string

	// en: A space-separated list of URLs to which, when the hyperlink is followed,
	// POST requests with the body PING will be sent by the browser (in the
	// background). Typically used for tracking.  For a more widely-supported feature
	// addressing the same use cases, see Navigator.sendBeacon().
	//     Value type: <list-of-URLs> ; Default value: none; Animatable: no
	Ping []string

	// en: Which referrer to send when fetching the URL.
	//     Value type: no-referrer|no-referrer-when-downgrade|same-origin|origin|
	//     strict-origin|origin-when-cross-origin|strict-origin-when-cross-origin|
	//     unsafe-url ; Default value: none; Animatable: no
	ReferrerPolicy int //todo: fazer

	// en: The relationship of the target object to the link object.
	//     Value type: <list-of-Link-Types> ; Default value: none; Animatable: yes
	Rel []string

	// en: Where to display the linked URL.
	//     Value type: _self|_parent|_top|_blank|<name> ; Default value: _self;
	//     Animatable: yes
	Target int //todo: fazer

	// en: A MIME type for the linked URL.
	//     Value type: <string> ; Default value: none; Animatable: yes
	Type string
}
