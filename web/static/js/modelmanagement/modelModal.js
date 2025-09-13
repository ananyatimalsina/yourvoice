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

function openModelModal(title, model = null, selected = false) {
	if (selected) {
		modelModalInputs = modelModal.querySelectorAll("input[unique='false']");
	} else {
		modelModalInputs = modelModal.querySelectorAll("input");
	}

	modelModalInputsUnique = modelModal.querySelectorAll("input[unique='true']");

	const messages = [];

	function handleUnique() {
		modelModalInputsUnique.forEach((input) => {
			if (input.parentElement.parentElement) {
				input.parentElement.parentElement.hidden = selected;
			}
		});
	}

	function resetForm() {
		modelModalInputs.forEach((input) => {
			input.value = "";
			messages.push("message-" + input.id);
		});
		if (selected) {
			modelModalTitle.textContent = `Edit ${title}`;
		} else modelModalTitle.textContent = `Add ${title}`;
	}

	function populateForm(model) {
		for (const key in model) {
			const input = modelModal.querySelector(`#${key}`);
			if (input) {
				input.value = model[key];
				messages.push("message-" + key);
			}
		}
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

	handleUnique();

	if (model) {
		populateForm(model);
		var method = "PUT",
			target = "row-" + model[ID_FIELD],
			swap = "update";
	} else {
		resetForm();
		if (selected) {
			var method = "PUT",
				targets = Array.from(selectedModels),
				swap = "update";
		} else {
			var method = "POST",
				target = "datatable",
				swap = "append";
		}
	}

	clearMessages();

	modelModalBackdrop.onclick = closeModelModal;
	modelModalClose.onclick = closeModelModal;
	modelModalSubmit.onclick = () => {
		const body = {};
		modelModalInputs.forEach((input) => {
			if (input.id) {
				body[input.id] = input.value;
			}
		});
		ajax("", {
			method: method,
			target: target,
			targets: targets,
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
