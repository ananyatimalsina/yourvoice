let modelModal = null;
let modelModalTitle = null;
let modelModalSubmit = null;
let modelModalClose = null;
let modelModalBackdrop = null;

function fetchElementsModelModal() {
	modelModal = document.getElementById("modelModal");
	modelModalTitle = document.getElementById("modelModalTitle");
	modelModalSubmit = document.getElementById("modelModalSubmit");
	modelModalClose = document.getElementById("modelModalClose");
	modelModalBackdrop = modelModal.querySelector("[data-tui-dialog-backdrop]");
}

document.addEventListener("DOMContentLoaded", fetchElementsModelModal);

function openModelModal(title, model = null) {
	let method = null;
	let target = null;
	let swap = null;
	if (model !== null) {
		method = "PUT";
		target = "row-" + model.ID;
		swap = "update";
		modelModalTitle.textContent = `Edit ${title}: ${model.name}`;

		// Add a hidden id input to the modelModal
		let idInput = modelModal.querySelector("#ID");
		if (!idInput) {
			idInput = document.createElement("input");
			idInput.type = "hidden";
			idInput.id = "ID";
			modelModal.appendChild(idInput);
		}
		idInput.value = model.ID;

		for (const key in model) {
			const input = modelModal.querySelector(`#${key}`);
			if (input) {
				input.value = model[key];
			}
		}
	} else {
		method = "POST";
		target = "datatable";
		swap = "append";
		modelModalTitle.textContent = `Add ${title}`;
		modelModal.querySelectorAll("input").forEach((input) => {
			input.value = "";
		});
		const idInput = modelModal.querySelector("#ID");
		if (idInput) {
			idInput.remove();
		}
	}

	function handleModalKeydown(e) {
		if (e.key === "Enter") {
			modelModalSubmit.click();
		}
	}

	function closeModal() {
		modelModal.removeEventListener("keydown", handleModalKeydown);
		modelModalBackdrop.onclick = null;
		modelModalClose.onclick = null;
		modelModalSubmit.onclick = null;
	}

	modelModalBackdrop.onclick = () => {
		closeModal();
	};

	modelModalClose.onclick = () => {
		closeModal();
	};

	modelModalSubmit.onclick = () => {
		const body = {};
		modelModal.querySelectorAll("input").forEach((input) => {
			if (input.id == "ID") {
				body[input.id] = parseInt(input.value);
			} else if (input.id) {
				body[input.id] = input.value;
			}
		});
		ajax("", { method: method, target: target, body: body, swap: swap })
			.then(() => {
				modelModalClose.click();
			})
			.catch((error) => {
				alert("Error: " + error.message);
			});
	};

	modelModal.addEventListener("keydown", handleModalKeydown);
}
