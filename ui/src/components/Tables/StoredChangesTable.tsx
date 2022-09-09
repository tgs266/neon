import React from "react"
import { Table, TableBody, TableCell, TableHead, TableRow, Tooltip } from '@mui/material';
import { Product } from "../../models/product";
import ErrorIcon from '@mui/icons-material/Error';
import { DateTime } from "luxon";
import { StoredChange } from "../../models/change";


export default function StoredChangesTable(props: { changes: StoredChange[] }) {
    const { changes } = props
    return <Table sx={{ minWidth: 650 }}>
        <TableHead>
            <TableRow>
                <TableCell>Product</TableCell>
                <TableCell align="right">Type</TableCell>
                <TableCell align="right">Version</TableCell>
                <TableCell align="right">Completed At</TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            {changes.map((row) => (
                <TableRow
                    key={row.id}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell component="th" scope="row">
                        {row.release.productName}
                    </TableCell>
                    <TableCell component="th" scope="row" align="right">
                        {row.type}
                    </TableCell>
                    <TableCell component="th" scope="row" align="right">
                        {row.release.productVersion}
                    </TableCell>
                    <TableCell component="th" scope="row" align="right">
                        {new Date(row.completedAt).toLocaleString()}
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table>
}