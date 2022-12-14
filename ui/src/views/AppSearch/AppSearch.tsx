import React, { useEffect, useState } from 'react'
import Grid from '@mui/material/Grid';
import { Button, Card, CardActions, CardContent, Table, TableBody, TableCell, TableHead, TableRow, TextField, Typography } from '@mui/material';
import { ProductService } from '../../services/ProductService';
import { AppService } from '../../services/AppService';
import { Product } from '../../models/product';
import { Link } from 'react-router-dom';
import { App } from '../../models/app';
import AddIcon from '@mui/icons-material/Add';
import Fab from '@mui/material/Fab';
import { CreateAppDialog } from './CreateAppDialog';

export function AppSearch() {

    const [name, setName] = useState<string>("")
    const [apps, setApps] = useState<App[]>([])
    const [open, setOpen] = useState(false)

    useEffect(() => {
        AppService.listApps(10, 0, name).then(r => {
            setApps(r.data.items)
        })
    }, [name])

    return <div>
        <Card sx={{ p: 1, display: "flex", alignItems: "center" }}>
            <TextField
                id="outlined-basic"
                label="Name"
                variant="outlined"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => setName(e.target.value)}
                value={name}
                fullWidth
            />
            <Fab sx={{ml: 1}} color="primary" aria-label="add" onClick={() => setOpen(true)}>
                <AddIcon />
            </Fab>
        </Card>

        <Card>
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell>Name</TableCell>
                        <TableCell align="right"></TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {apps.map((row) => (
                        <TableRow
                            key={row.name}
                            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                            <TableCell component="th" scope="row">
                                {row.name}
                            </TableCell>
                            <TableCell align="right"><Button component={Link} to={`/apps/${row.name}`}>View</Button></TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </Card>
        <CreateAppDialog open={open} setOpen={setOpen} />
    </div>
}