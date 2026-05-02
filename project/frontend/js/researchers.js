(function () {
    const {
        $,
        apiRequest,
        formToObject,
        setFormValues,
        resetForm,
        renderHead,
        renderTable,
        filterRows,
        populateSelect,
        showStatus,
        clearStatus,
        toNumber
    } = window.RPT;

    const form = $("#researcherForm");
    const idInput = $("#researcherId");
    const formTitle = $("#formTitle");
    const formStatus = $("#formStatus");
    const tableStatus = $("#tableStatus");
    const searchInput = $("#searchInput");
    const departmentSelect = $("#departmentId");
    const columns = [
        { key: "researcher_id", label: "ID" },
        { key: "full_name", label: "Full Name" },
        { key: "email", label: "Email" },
        { key: "role", label: "Role" },
        { key: "department_name", label: "Department" }
    ];

    let researchers = [];

    function actionButtons(row) {
        return `
            <div class="row-actions">
                <button class="compact secondary" data-action="edit" data-id="${row.researcher_id}">Edit</button>
                <button class="compact danger" data-action="delete" data-id="${row.researcher_id}">Delete</button>
            </div>
        `;
    }

    function drawTable() {
        const visibleRows = filterRows(researchers, searchInput.value, ["full_name", "email", "role", "department_name"]);
        renderTable($("#researchersBody"), visibleRows, columns, actionButtons);
    }

    async function loadLookups() {
        const departments = await apiRequest("/api/departments");
        populateSelect(departmentSelect, departments, "department_id", "department_name", "Select department");
    }

    async function loadResearchers() {
        clearStatus(tableStatus);
        try {
            researchers = await apiRequest("/api/researchers");
            drawTable();
        } catch (error) {
            showStatus(tableStatus, error.message, true);
        }
    }

    function resetResearcherForm() {
        resetForm(form);
        formTitle.textContent = "Add Researcher";
        $("#saveButton").textContent = "Save";
        clearStatus(formStatus);
    }

    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        clearStatus(formStatus);

        const payload = formToObject(form);
        payload.department_id = toNumber(payload.department_id);

        const id = idInput.value;
        const path = id ? `/api/researchers/${id}` : "/api/researchers";
        const method = id ? "PUT" : "POST";

        try {
            await apiRequest(path, {
                method,
                body: JSON.stringify(payload)
            });
            showStatus(formStatus, id ? "Researcher updated." : "Researcher added.");
            resetResearcherForm();
            await loadResearchers();
        } catch (error) {
            showStatus(formStatus, error.message, true);
        }
    });

    $("#resetButton").addEventListener("click", resetResearcherForm);
    searchInput.addEventListener("input", drawTable);

    $("#researchersBody").addEventListener("click", async (event) => {
        const button = event.target.closest("button");
        if (!button) {
            return;
        }

        const id = Number(button.dataset.id);
        const researcher = researchers.find((item) => item.researcher_id === id);
        if (!researcher) {
            return;
        }

        if (button.dataset.action === "edit") {
            setFormValues(form, researcher);
            idInput.value = researcher.researcher_id;
            departmentSelect.value = researcher.department_id;
            formTitle.textContent = "Edit Researcher";
            $("#saveButton").textContent = "Update";
            clearStatus(formStatus);
            return;
        }

        if (button.dataset.action === "delete" && confirm("Delete this researcher?")) {
            try {
                await apiRequest(`/api/researchers/${id}`, { method: "DELETE" });
                showStatus(tableStatus, "Researcher deleted.");
                await loadResearchers();
                resetResearcherForm();
            } catch (error) {
                showStatus(tableStatus, error.message, true);
            }
        }
    });

    renderHead($("#researchersHead"), columns, true);
    loadLookups()
        .then(loadResearchers)
        .catch((error) => showStatus(tableStatus, error.message, true));
})();
