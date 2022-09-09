import React from "react"
import { Table, TableBody, TableCell, TableHead, TableRow, Tooltip } from '@mui/material';
import { Product } from "../../models/product";
import ErrorIcon from '@mui/icons-material/Error';
import { Link } from "react-router-dom";
import { App } from "../../models/app";


export default function AppInstallsTable(props: { app?: App }) {
    const { app } = props
    return <Table sx={{ minWidth: 650 }}>
        <TableHead>
            <TableRow>
                <TableCell>Product Name</TableCell>
                <TableCell align="right">Version</TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            {app?.installs && app.installs.map((row) => (
                <TableRow
                    key={row.productName}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell component="th" scope="row">
                        <Link to={`/apps/${app.name}/installs/${row.productName}`}>
                            {row.productName}
                        </Link>
                    </TableCell>
                    <TableCell align="right" component="th" scope="row">
                        {row.releaseVersion}
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table>
}