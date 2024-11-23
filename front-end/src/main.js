import * as VTable from '@visactor/vtable';

async function fetchData() {
    try {
        const response = await fetch('http://localhost:8831/api/data', {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        console.log('Fetched data:', data);
        return data;
    } catch (error) {
        console.error('Fetch error:', error);
        return [];
    }
}

async function initTable() {
    try {
        const records = await fetchData();
        console.log('Records for table:', records);

        const columns = [
            { field: 'id', title: 'ID', width: 80 },
            { field: 'name', title: '姓名', width: 120 },
            { field: 'age', title: '年龄', width: 100 },
            { field: 'city', title: '城市', width: 120 }
        ];

        const option = {
            records,
            columns,
            widthMode: 'standard',
            width: 1200,
            height: 600
        };

        const container = document.getElementById('tableContainer');
        if (!container) {
            console.error('Table container not found!');
            return;
        }

        const tableInstance = new VTable.ListTable(container, option);
        console.log('Table instance created');

        setInterval(async () => {
            try {
                const newData = await fetchData();
                if (Array.isArray(newData) && newData.length > 0) {
                    console.log('Updating table with new data:', newData);
                    tableInstance.setRecords(newData);
                }
            } catch (error) {
                console.error('Error updating table:', error);
            }
        }, 5000);
    } catch (error) {
        console.error('Error initializing table:', error);
    }
}

document.addEventListener('DOMContentLoaded', initTable);