let modelBar = null;
let searchInput = null;
let table = null;

function fetchElementsFilterBar() {
	modelBar = document.getElementById("modelBar");
	searchInput = document.getElementById("search");
	searchInput.addEventListener("input", debounce(performSearch, 400));
	table = document.getElementById("datatable").parentElement;
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

function toggleVisibility(element, colIndex) {
	const isHidden = table.rows[0]?.cells[colIndex]?.style.display === "none";
	for (var row of table.rows) {
		if (row.cells.length > colIndex) {
			row.cells[colIndex].style.display = isHidden ? "" : "none";
		}
	}
	element.firstElementChild.firstElementChild.setAttribute(
		"visibility",
		isHidden ? "visible" : "hidden",
	);
}
