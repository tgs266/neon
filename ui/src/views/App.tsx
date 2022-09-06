import React, { useEffect, useState } from 'react'
import Grid from '@mui/material/Grid';
import { Badge, Button, Card, CardActions, CardContent, Chip, Table, TableBody, TableCell, TableHead, TableRow, TextField, Tooltip, Typography } from '@mui/material';
import { ProductService } from '../services/ProductService';
import { AppService } from '../services/AppService';
import { Product } from '../models/product';
import { Link } from 'react-router-dom';
import { useParams } from 'react-router';
import { App } from '../models/app';
import { Box } from '@mui/system';
import ErrorIcon from '@mui/icons-material/Error';
import TitleCard from '../components/DetailsTitleCard';
import Statistic from '../components/Statistic';

export function App() {

    const [app, setApp] = useState<App>(null)
    const { name } = useParams();

    useEffect(() => {
        AppService.get(name).then(r => {
            setApp(r.data)
        })
    }, [name])

    return <div>
        <TitleCard title={<div style={{ display: "flex", alignItems: "center" }}>
            {app?.name ? <div>{app.name}</div> : null}
            <Chip sx={{ ml: 1 }} label={app?.installStatus} />
        </div>} sx={{ mb: 1 }}>
            <Box sx={{ mt: 1, display: "flex", alignItems: "center", gap: 2 }}>
                <Statistic label="Created At" value={new Date(app?.createdAt).toLocaleString()} />
                <Statistic label="Updated At" value={new Date(app?.updatedAt).toLocaleString()} />
            </Box>
        </TitleCard>

        <Card>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell>Product</TableCell>
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
                                <div style={{ display: "flex", alignItems: "center" }}>
                                    {row.productName}
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
        </Card>
    </div>
}