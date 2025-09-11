let modelBar = null;
let searchInput = null;

function fetchElementsFilterBar() {
	modelBar = document.getElementById("modelBar");
	searchInput = document.getElementById("search");
	searchInput.addEventListener("input", debounce(performSearch, 400));
}

document.addEventListener("DOMContentLoaded", fetchElementsFilterBar);

function debounce(func, delay) {
	let timeoutId;
	return function (...args) {
		clearTimeout(timeoutId);
		timeoutId = setTimeout(() => {
			func.apply(this, args);
		}, delay);
	};
}

function performSearch() {
	const params = new URLSearchParams();
	const inputs = modelBar.querySelectorAll("input, select");
	inputs.forEach((input) => {
		if (input.value) {
			params.append(input.name, input.value);
		}
	});
	ajax("?" + params, { target: "datatable", swap: "update" });
}
