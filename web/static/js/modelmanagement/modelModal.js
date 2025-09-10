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

const ID_FIELD = "ID";

function openModelModal(title, model = null) {
	const modelModalInputs = Array.from(modelModal.querySelectorAll("input"));
	const messages = [];

	function resetForm() {
		modelModalInputs.forEach((input) => {
			input.value = "";
			messages.push("message-" + input.id);
		});
		const idInput = modelModal.querySelector(`#${ID_FIELD}`);
		if (idInput) idInput.remove();
		modelModalTitle.textContent = `Add ${title}`;
	}

	function populateForm(model) {
		for (const key in model) {
			if (key === ID_FIELD) continue;
			const input = modelModal.querySelector(`#${key}`);
			if (input) {
				input.value = model[key];
				messages.push("message-" + key);
			}
		}
		let idInput = modelModal.querySelector(`#${ID_FIELD}`);
		if (!idInput) {
			idInput = document.createElement("input");
			idInput.type = "hidden";
			idInput.id = ID_FIELD;
			modelModal.appendChild(idInput);
			modelModalInputs.push(idInput);
		}
		idInput.value = model[ID_FIELD];
		modelModalTitle.textContent = `Edit ${title}: ${model.name}`;
	}

	function clearMessages() {
		for (const messageId of messages) {
			const message = document.getElementById(messageId);
			if (message) message.textContent = "";
		}
	}

	function closeModelModal() {
		modelModal.removeEventListener("keydown", handleModelModalKeydown);
		modelModalBackdrop.onclick = null;
		modelModalClose.onclick = null;
		modelModalSubmit.onclick = null;
	}

	function handleModelModalKeydown(e) {
		if (e.key === "Enter") {
			modelModalSubmit.click();
		}
	}

	if (model) {
		populateForm(model);
		var method = "PUT",
			target = "row-" + model[ID_FIELD],
			swap = "update";
	} else {
		resetForm();
		var method = "POST",
			target = "datatable",
			swap = "append";
	}

	clearMessages();

	modelModalBackdrop.onclick = closeModelModal;
	modelModalClose.onclick = closeModelModal;
	modelModalSubmit.onclick = () => {
		const body = {};
		modelModalInputs.forEach((input) => {
			if (input.id) {
				body[input.id] =
					input.id === ID_FIELD ? parseInt(input.value, 10) : input.value;
			}
		});
		ajax("", {
			method: method,
			target: target,
			body: body,
			swap: swap,
		})
			.then(() => modelModalClose.click())
			.catch((error) => {
				if (error.status === 400) {
					messages.forEach((messageId) => {
						const message = document.getElementById(messageId);
						message.textContent = error.message[messageId] || "";
					});
				} else {
					alert("Error: " + error.message);
				}
			});
	};

	modelModal.addEventListener("keydown", handleModelModalKeydown);
}
