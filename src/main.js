// 修改导入方式
import * as VTable from '@visactor/vtable';
// import '@visactor/vtable/dist/index.css';
import { records, columns } from './data.js';

document.addEventListener('DOMContentLoaded', () => {
    const option = {
        records,
        columns,
        widthMode: 'standard',
        width: 1200,
        height: 600
    };

    try {
        // 使用 VTable.ListTable 而不是直接使用 ListTable
        const tableInstance = new VTable.ListTable(
            document.getElementById('tableContainer'), 
            option
        );
        console.log('Table created successfully:', tableInstance);
    } catch (error) {
        console.error('Error creating table:', error);
    }
});