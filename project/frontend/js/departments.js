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
        showStatus,
        clearStatus
    } = window.RPT;

    const form = $("#departmentForm");
    const idInput = $("#departmentId");
    const formTitle = $("#formTitle");
    const formStatus = $("#formStatus");
    const tableStatus = $("#tableStatus");
    const searchInput = $("#searchInput");
    const columns = [
        { key: "department_id", label: "ID" },
        { key: "department_name", label: "Department" },
        { key: "college", label: "College" }
    ];

    let departments = [];

    function actionButtons(row) {
        return `
            <div class="row-actions">
                <button class="compact secondary" data-action="edit" data-id="${row.department_id}">Edit</button>
                <button class="compact danger" data-action="delete" data-id="${row.department_id}">Delete</button>
            </div>
        `;
    }

    function drawTable() {
        const visibleRows = filterRows(departments, searchInput.value, ["department_name", "college"]);
        renderTable($("#departmentsBody"), visibleRows, columns, actionButtons);
    }

    async function loadDepartments() {
        clearStatus(tableStatus);
        try {
            departments = await apiRequest("/api/departments");
            drawTable();
        } catch (error) {
            showStatus(tableStatus, error.message, true);
        }
    }

    function resetDepartmentForm() {
        resetForm(form);
        formTitle.textContent = "Add Department";
        $("#saveButton").textContent = "Save";
        clearStatus(formStatus);
    }

    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        clearStatus(formStatus);

        const payload = formToObject(form);
        const id = idInput.value;
        const path = id ? `/api/departments/${id}` : "/api/departments";
        const method = id ? "PUT" : "POST";

        try {
            await apiRequest(path, {
                method,
                body: JSON.stringify(payload)
            });
            showStatus(formStatus, id ? "Department updated." : "Department added.");
            resetDepartmentForm();
            await loadDepartments();
        } catch (error) {
            showStatus(formStatus, error.message, true);
        }
    });

    $("#resetButton").addEventListener("click", resetDepartmentForm);
    searchInput.addEventListener("input", drawTable);

    $("#departmentsBody").addEventListener("click", async (event) => {
        const button = event.target.closest("button");
        if (!button) {
            return;
        }

        const id = Number(button.dataset.id);
        const department = departments.find((item) => item.department_id === id);
        if (!department) {
            return;
        }

        if (button.dataset.action === "edit") {
            setFormValues(form, department);
            idInput.value = department.department_id;
            formTitle.textContent = "Edit Department";
            $("#saveButton").textContent = "Update";
            clearStatus(formStatus);
            return;
        }

        if (button.dataset.action === "delete" && confirm("Delete this department?")) {
            try {
                await apiRequest(`/api/departments/${id}`, { method: "DELETE" });
                showStatus(tableStatus, "Department deleted.");
                await loadDepartments();
                resetDepartmentForm();
            } catch (error) {
                showStatus(tableStatus, error.message, true);
            }
        }
    });

    renderHead($("#departmentsHead"), columns, true);
    loadDepartments();
})();
