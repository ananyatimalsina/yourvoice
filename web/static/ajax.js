async function ajax(url, params = {}) {
	const response = await fetch(url, {
		method: params.method || "GET",
		headers: {
			"AJAX-Target": params.target || "",
			"AJAX-Targets": params.targets || "",
		},
		body: JSON.stringify(params.body),
	});

	if (!response.ok) {
		const text = await response.text();
		let json = null;
		try {
			json = JSON.parse(text);
		} catch (e) {
			throw {
				status: response.status,
				message: text,
			};
		}

		throw {
			status: response.status,
			message: json,
		};
	}

	// TODO: look into DomPurify for HTML sanitization, but with hooks to allow calling self
	const body = await response.text();
	const parser = new DOMParser();
	const responseDoc = parser.parseFromString(body, "text/html");

	// Helper to do swap
	function swapElements(target, swapType, responseElement = null) {
		switch (swapType) {
			case "update":
				target.innerHTML = responseElement ? responseElement.innerHTML : body;
				break;
			case "before":
				target.insertAdjacentHTML(
					"beforebegin",
					responseElement ? responseElement.outerHTML : body,
				);
				break;
			case "after":
				target.insertAdjacentHTML(
					"afterend",
					responseElement ? responseElement.outerHTML : body,
				);
				break;
			case "prepend":
				target.insertAdjacentHTML(
					"afterbegin",
					responseElement ? responseElement.outerHTML : body,
				);
				break;
			case "append":
				target.insertAdjacentHTML(
					"beforeend",
					responseElement ? responseElement.outerHTML : body,
				);
				break;
			case "remove":
				target.remove();
				break;
			default:
				target.outerHTML = responseElement ? responseElement.outerHTML : body;
				break;
		}
	}

	if (params.targets) {
		for (const targetId of params.targets) {
			const target = document.getElementById(targetId);
			const responseElement = responseDoc.getElementById(targetId);
			swapElements(target, params.swap, responseElement);
		}
	} else {
		const target = document.getElementById(params.target);
		swapElements(target, params.swap);
	}

	if (params.pushState !== false) {
		history.pushState({}, "", url);
	}

	if (params.swap == "remove") {
		if (params.targets) {
			selectedModels.clear();
		} else {
			selectedModels.delete(params.target.slice(4));
		}
		updateUIState();
	} else {
		if (params.modelManagementRequest === true) {
			selectedModels.clear();
			updatePageLoad();
		} else {
			updateUIState();
		}
	}
}

function updatePageLoad() {
	updateSelectedModels();
	fetchElementsModelModal();
}
