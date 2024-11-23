import { ListTable } from '@visactor/vtable';

async function fetchData() {
    try {
        const response = await fetch('/api/data');
        return await response.json();
    } catch (error) {
        console.error('Error fetching data:', error);
        return [];
    }
}

async function initTable() {
    const records = await fetchData();

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

    const tableInstance = new ListTable(
        document.getElementById('tableContainer'),
        option
    );

    // 定期刷新数据
    setInterval(async () => {
        const newData = await fetchData();
        tableInstance.updateData(newData);
    }, 5000);
}

document.addEventListener('DOMContentLoaded', initTable);