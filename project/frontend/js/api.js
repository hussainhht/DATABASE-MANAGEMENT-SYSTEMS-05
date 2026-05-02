(function () {
    const API_BASE = "";

    function $(selector) {
        return document.querySelector(selector);
    }

    function $all(selector) {
        return Array.from(document.querySelectorAll(selector));
    }

    function escapeHTML(value) {
        if (value === null || value === undefined) {
            return "";
        }

        return String(value)
            .replaceAll("&", "&amp;")
            .replaceAll("<", "&lt;")
            .replaceAll(">", "&gt;")
            .replaceAll('"', "&quot;")
            .replaceAll("'", "&#039;");
    }

    async function apiRequest(path, options = {}) {
        const response = await fetch(API_BASE + path, {
            headers: {
                "Content-Type": "application/json",
                ...(options.headers || {})
            },
            ...options
        });

        let data = null;
        const text = await response.text();
        if (text) {
            data = JSON.parse(text);
        }

        if (!response.ok) {
            const message = data && data.error ? data.error : "Request failed";
            throw new Error(message);
        }

        return data;
    }

    function showStatus(element, message, isError = false) {
        if (!element) {
            return;
        }

        element.textContent = message;
        element.className = isError ? "status error" : "status success";
    }

    function clearStatus(element) {
        if (!element) {
            return;
        }

        element.textContent = "";
        element.className = "status";
    }

    function formToObject(form) {
        const data = {};
        const formData = new FormData(form);

        for (const [key, value] of formData.entries()) {
            data[key] = value;
        }

        return data;
    }

    function setFormValues(form, data) {
        Object.keys(data).forEach((key) => {
            const input = form.elements[key];
            if (input) {
                input.value = data[key] ?? "";
            }
        });
    }

    function resetForm(form) {
        form.reset();
        const idInput = form.querySelector("[data-record-id]");
        if (idInput) {
            idInput.value = "";
        }
    }

    function renderTable(tbody, rows, columns, actions) {
        if (!tbody) {
            return;
        }

        if (!rows.length) {
            tbody.innerHTML = `<tr><td colspan="${columns.length + (actions ? 1 : 0)}" class="empty-state">No records found.</td></tr>`;
            return;
        }

        tbody.innerHTML = rows.map((row) => {
            const cells = columns.map((column) => `<td>${escapeHTML(row[column.key])}</td>`).join("");
            const actionCell = actions ? `<td>${actions(row)}</td>` : "";
            return `<tr>${cells}${actionCell}</tr>`;
        }).join("");
    }

    function renderHead(thead, columns, hasActions) {
        if (!thead) {
            return;
        }

        const actionHead = hasActions ? "<th>Actions</th>" : "";
        thead.innerHTML = `<tr>${columns.map((column) => `<th>${escapeHTML(column.label)}</th>`).join("")}${actionHead}</tr>`;
    }

    function filterRows(rows, query, fields) {
        const text = query.trim().toLowerCase();
        if (!text) {
            return rows;
        }

        return rows.filter((row) => fields.some((field) => String(row[field] ?? "").toLowerCase().includes(text)));
    }

    function populateSelect(select, rows, valueKey, labelKey, placeholder) {
        if (!select) {
            return;
        }

        select.innerHTML = `<option value="">${escapeHTML(placeholder)}</option>` +
            rows.map((row) => `<option value="${escapeHTML(row[valueKey])}">${escapeHTML(row[labelKey])}</option>`).join("");
    }

    function toNumber(value) {
        const numberValue = Number(value);
        return Number.isNaN(numberValue) ? 0 : numberValue;
    }

    window.RPT = {
        $,
        $all,
        escapeHTML,
        apiRequest,
        showStatus,
        clearStatus,
        formToObject,
        setFormValues,
        resetForm,
        renderTable,
        renderHead,
        filterRows,
        populateSelect,
        toNumber
    };
})();
