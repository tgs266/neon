import React, { useEffect, useState } from 'react'
import Grid from '@mui/material/Grid';
import { Badge, Button, Card, CardActions, CardContent, Chip, Tab, Table, TableBody, TableCell, TableHead, TableRow, Tabs, TextField, Tooltip, Typography } from '@mui/material';
import { ProductService } from '../../services/ProductService';
import { AppService } from '../../services/AppService';
import { Product } from '../../models/product';
import { Link, useSearchParams } from 'react-router-dom';
import { useParams } from 'react-router';
import { App } from '../../models/app';
import { Box } from '@mui/system';
import ErrorIcon from '@mui/icons-material/Error';
import TitleCard from '../../components/DetailsTitleCard';
import Statistic from '../../components/Statistic';
import InstallsTable from '../../components/Tables/InstallsTable';
import AppInstallsTable from '../../components/Tables/AppInstallTable';
import ChangesTable from '../../components/Tables/StoredChangesTable';
import { QueuedChange, StoredChange } from '../../models/change';
import QueuedChangesTable from '../../components/Tables/QueuedChangesTable';
import StoredChangesTable from '../../components/Tables/StoredChangesTable';
import { AddProductDialog } from './AddProductDialog';

function a11yProps(index: number) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}

export function App() {

    const [open, setOpen] = useState(false)
    const [app, setApp] = useState<App>(null)
    const [storedChanges, setStoredChanges] = useState<StoredChange[]>([])
    const [queuedChanges, setQueuedChanges] = useState<QueuedChange[]>([])
    const { name } = useParams();
    const [search, setSearch] = useSearchParams();


    useEffect(() => {
        AppService.get(name).then(r => {
            setApp(r.data)
        })
        AppService.getStoredChanges(name).then(r => {
            setStoredChanges(r.data.items)
        })
        AppService.getQueuedChanges(name).then(r => {
            setQueuedChanges(r.data.items)
        })
    }, [name])

    if (search.get("tab") == null) {
        setSearch({tab: "0"}, {replace: true})
    }

    const handleChange = (event: React.SyntheticEvent, newValue: number) => {
        setSearch({tab: newValue.toString()}, {replace: true})
    };

    return <div>
        <TitleCard title={<div style={{ display: "flex", alignItems: "center" }}>
            {app?.name ? <div>{app.name}</div> : null}
        </div>}>
            <Box sx={{ mt: 1, display: "flex", alignItems: "center", gap: 2 }}>
                <Statistic label="Created At" value={new Date(app?.createdAt).toLocaleString()} />
                <Statistic label="Updated At" value={new Date(app?.updatedAt).toLocaleString()} />
            </Box>
            <Button onClick={() => setOpen(true)}>Add Product</Button>
            <AddProductDialog appName={app.name} open={open} setOpen={setOpen} />
        </TitleCard>

        <Box sx={{}}>
            <Tabs sx={{ml: 1}} value={parseInt(search.get("tab"))} onChange={handleChange} aria-label="basic tabs example">
                <Tab label="Installs" {...a11yProps(0)} />
                <Tab label="Changes" {...a11yProps(1)} />
                <Tab label="Queued Changes" {...a11yProps(2)} />
            </Tabs>
        </Box>

        {search.get("tab") === "0" &&
            <Card>
                <AppInstallsTable app={app} />
            </Card>
        }
        {search.get("tab") === "1" &&
            <Card>
                <StoredChangesTable changes={storedChanges} />
            </Card>
        }
        {search.get("tab") === "2" &&
            <Card>
                <QueuedChangesTable changes={queuedChanges} />
            </Card>
        }
    </div>
}