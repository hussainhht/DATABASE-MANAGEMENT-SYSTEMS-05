(function () {
    const {
        $,
        apiRequest,
        renderHead,
        renderTable,
        showStatus,
        clearStatus
    } = window.RPT;

    const reports = {
        authors: {
            title: "Publications with Authors",
            path: "/api/reports/publications-with-authors",
            columns: [
                { key: "publication_id", label: "Publication ID" },
                { key: "title", label: "Title" },
                { key: "publication_year", label: "Year" },
                { key: "publication_type", label: "Type" },
                { key: "venue_name", label: "Venue" },
                { key: "researcher_name", label: "Researcher" },
                { key: "author_order", label: "Order" }
            ]
        },
        departments: {
            title: "Publications by Department",
            path: "/api/reports/publications-by-department",
            columns: [
                { key: "department_id", label: "Department ID" },
                { key: "department_name", label: "Department" },
                { key: "publication_count", label: "Publications" }
            ]
        },
        researchers: {
            title: "Publications by Researcher",
            path: "/api/reports/publications-by-researcher",
            columns: [
                { key: "researcher_id", label: "Researcher ID" },
                { key: "full_name", label: "Full Name" },
                { key: "publication_count", label: "Publications" }
            ]
        },
        years: {
            title: "Publications by Year",
            path: "/api/reports/publications-by-year",
            columns: [
                { key: "publication_year", label: "Year" },
                { key: "publication_count", label: "Publications" }
            ]
        },
        moreThanOne: {
            title: "Researchers with More Than One Publication",
            path: "/api/reports/researchers-more-than-one-publication",
            columns: [
                { key: "researcher_id", label: "Researcher ID" },
                { key: "full_name", label: "Full Name" },
                { key: "publication_count", label: "Publications" }
            ]
        },
        keyword: {
            title: "Publications by Keyword",
            path: "/api/reports/publications-by-keyword",
            parameterName: "keyword",
            parameterLabel: "Keyword",
            defaultValue: "Machine",
            columns: [
                { key: "publication_id", label: "Publication ID" },
                { key: "title", label: "Title" },
                { key: "keyword_text", label: "Keyword" }
            ]
        },
        type: {
            title: "Publications by Type",
            path: "/api/reports/publications-by-type",
            parameterName: "type",
            parameterLabel: "Publication Type",
            defaultValue: "Journal",
            columns: [
                { key: "publication_id", label: "Publication ID" },
                { key: "title", label: "Title" },
                { key: "publication_year", label: "Year" },
                { key: "publication_type", label: "Type" }
            ]
        },
        venue: {
            title: "Publications by Venue",
            path: "/api/reports/publications-by-venue",
            columns: [
                { key: "publication_id", label: "Publication ID" },
                { key: "title", label: "Title" },
                { key: "venue_name", label: "Venue" },
                { key: "venue_type", label: "Venue Type" },
                { key: "publisher", label: "Publisher" }
            ]
        },
        latest: {
            title: "Latest Publications",
            path: "/api/reports/latest-publications",
            columns: [
                { key: "publication_id", label: "Publication ID" },
                { key: "title", label: "Title" },
                { key: "publication_year", label: "Year" },
                { key: "publication_type", label: "Type" }
            ]
        },
        search: {
            title: "Search Publications by Title",
            path: "/api/reports/search-publications",
            parameterName: "title",
            parameterLabel: "Title",
            defaultValue: "energy",
            columns: [
                { key: "publication_id", label: "Publication ID" },
                { key: "title", label: "Title" },
                { key: "publication_year", label: "Year" },
                { key: "publication_type", label: "Type" }
            ]
        }
    };

    const reportSelect = $("#reportSelect");
    const parameterRow = $("#parameterRow");
    const parameterInput = $("#parameterInput");
    const parameterLabel = $("#parameterLabel");
    const status = $("#reportStatus");

    function loadReportOptions() {
        reportSelect.innerHTML = Object.entries(reports)
            .map(([key, report]) => `<option value="${key}">${report.title}</option>`)
            .join("");
    }

    function selectedReport() {
        return reports[reportSelect.value];
    }

    function updateParameterControl() {
        const report = selectedReport();
        if (report.parameterName) {
            parameterRow.style.display = "grid";
            parameterLabel.textContent = report.parameterLabel;
            parameterInput.value = report.defaultValue;
        } else {
            parameterRow.style.display = "none";
            parameterInput.value = "";
        }
    }

    function buildPath(report) {
        if (!report.parameterName) {
            return report.path;
        }

        const value = parameterInput.value.trim();
        if (!value) {
            throw new Error(`${report.parameterLabel} is required`);
        }

        const params = new URLSearchParams({ [report.parameterName]: value });
        return `${report.path}?${params.toString()}`;
    }

    async function runReport() {
        clearStatus(status);
        const report = selectedReport();
        $("#reportTitle").textContent = report.title;
        renderHead($("#reportHead"), report.columns, false);

        try {
            const rows = await apiRequest(buildPath(report));
            renderTable($("#reportBody"), rows, report.columns);
            $("#rowCount").textContent = `${rows.length} rows`;
        } catch (error) {
            $("#rowCount").textContent = "0 rows";
            renderTable($("#reportBody"), [], report.columns);
            showStatus(status, error.message, true);
        }
    }

    reportSelect.addEventListener("change", () => {
        updateParameterControl();
        runReport();
    });

    $("#runReportButton").addEventListener("click", runReport);

    loadReportOptions();
    updateParameterControl();
    runReport();
})();
