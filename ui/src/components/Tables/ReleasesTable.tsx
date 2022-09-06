import React from "react"
import { Table, TableBody, TableCell, TableHead, TableRow } from '@mui/material';
import { Product } from "../../models/product";


export default function ReleasesTable(props: { product?: Product }) {
    const { product } = props
    return <Table sx={{ minWidth: 650 }}>
        <TableHead>
            <TableRow>
                <TableCell>Version</TableCell>
                <TableCell align="right">Release Channel</TableCell>
                <TableCell align="right">Helm Chart</TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            {product?.releases && product.releases.map((row) => (
                <TableRow
                    key={row.productVersion}
                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                >
                    <TableCell component="th" scope="row">
                        {row.productVersion}
                    </TableCell>
                    <TableCell align="right" component="th" scope="row">
                        {row.releaseChannel}
                    </TableCell>
                    <TableCell align="right" component="th" scope="row">
                        {row.helmChart}
                    </TableCell>
                </TableRow>
            ))}
        </TableBody>
    </Table>
}