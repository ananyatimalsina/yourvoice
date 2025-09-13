let searchParams = new URLSearchParams(window.location.search);

let modelBar = null;
let searchInput = null;
let table = null;

function fetchElementsFilterBar() {
	modelBar = document.getElementById("modelBar");
	searchInput = document.getElementById("search");
	searchInput.addEventListener("input", debounce(performSearch, 500));
	table = document.getElementById("datatable").parentElement;
}

document.addEventListener("DOMContentLoaded", fetchElementsFilterBar);

function debounce(func, delay) {
	let timeoutId;
	return () => {
		clearTimeout(timeoutId);
		timeoutId = setTimeout(() => {
			func.apply();
		}, delay);
	};
}

function compareSearchParams(params1, params2) {
	// Convert params to sorted arrays for comparison
	const arr1 = Array.from(params1.entries()).sort();
	const arr2 = Array.from(params2.entries()).sort();

	// Compare length first
	if (arr1.length !== arr2.length) return false;

	// Compare each key/value pair
	for (let i = 0; i < arr1.length; i++) {
		if (arr1[i][0] !== arr2[i][0] || arr1[i][1] !== arr2[i][1]) {
			return false;
		}
	}
	return true;
}

function performSearch(appendParams = null) {
	let combined = new URLSearchParams();
	const inputs = modelBar.querySelectorAll("input, select");
	inputs.forEach((input) => {
		combined.set(input.name, input.value);
	});
	if (appendParams) {
		combined = new URLSearchParams({
			...Object.fromEntries(combined),
			...Object.fromEntries(new URLSearchParams(appendParams)),
		});
	}
	if (!compareSearchParams(searchParams, combined)) {
		ajax("?" + combined, { target: "datatable", swap: "update" }).then(() => {
			searchParams = combined;
		});
	}
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
