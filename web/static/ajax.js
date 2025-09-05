async function ajax(url, params = {}) {
	const response = await fetch(url, {
		method: params.method || "GET",
		headers: {
			"AJAX-Request": "true",
			"AJAX-Target": params.target || "",
		},
		body: params.body,
	});

	if (!response.ok) {
		const text = await response.text();
		throw new Error(
			`Request failed: ${response.status} ${response.statusText}\n${text}`,
		);
	}

	// TODO: look into DomPurify for HTML sanitization, but with hooks to allow calling self

	target = document.getElementById(params.target);
	body = await response.text();

	switch (params.swap) {
		case "update":
			target.innerHTML = body;
			break;
		case "before":
			target.insertAdjacentHTML("beforebegin", body);
			break;
		case "after":
			target.insertAdjacentHTML("afterend", body);
			break;
		case "prepend":
			target.insertAdjacentHTML("afterbegin", body);
			break;
		case "append":
			target.insertAdjacentHTML("beforeend", body);
			break;
		case "remove":
			target.remove();
			break;
		default:
			target.outerHTML = body;
			break;
	}

	if (params.pushState !== false) {
		history.pushState({}, "", url);
	}
}
