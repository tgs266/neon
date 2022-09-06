import React from "react"
import { Table, TableBody, TableCell, TableHead, TableRow, Tooltip } from '@mui/material';
import { Product } from "../../models/product";
import ErrorIcon from '@mui/icons-material/Error';
import { Link } from "react-router-dom";


export default function InstallsTable(props: { product?: Product }) {
    const { product } = props
    return <Table sx={{ minWidth: 650 }}>
        <TableHead>
            <TableRow>
                <TableCell>App</TableCell>
                <TableCell align="right">Version</TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            {product?.installs && product.installs.map((row) => (
                <TableRow
                    key={row.appName}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell component="th" scope="row">
                        <div style={{ display: "flex", alignItems: "center" }}>
                            <Link to={`/apps/${row.appName}`}>
                                {row.appName}
                            </Link>
                            {row.error &&
                                <Tooltip title={row.error}>
                                    <ErrorIcon color="secondary" />
                                </Tooltip>
                            }
                        </div>
                    </TableCell>
                    <TableCell align="right" component="th" scope="row">
                        {row.releaseVersion}
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table>
}