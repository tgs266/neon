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
import { Install } from '../../models/install';
import { ResourceList } from '../../models/resources';
import ResourceContainer from '../../components/Resources/ResourceContainer';

function a11yProps(index: number) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}

export function Install() {

    const [install, setInstall] = useState<Install>(null)
    const [resources, setResources] = useState<ResourceList>(null)
    const { name, productName } = useParams();


    useEffect(() => {
        AppService.getInstall(name, productName).then(r => {
            setInstall(r.data)
        })
        AppService.getInstallResources(name, productName).then(r => {
            setResources(r.data)
        })
    }, [name])

    return <div style={{display: "flex", flexDirection: "column", height: "calc(100%)"}}>
        <TitleCard title={<div style={{ display: "flex", alignItems: "center" }}>
            <div>{name}/{productName}</div>
        </div>} sx={{ flexShrink: 0, zIndex: 100, mb: 0 }}>
            <Box sx={{ mt: 1, display: "flex", alignItems: "center", gap: 2 }}>
                <Statistic label="Created At" value={new Date(install?.createdAt).toLocaleString()} />
                <Statistic label="Updated At" value={new Date(install?.updatedAt).toLocaleString()} />
            </Box>
        </TitleCard>

        <Box sx={{ flexGrow: 1, height: "100%", overflowY: "auto" }}>
            <ResourceContainer appName={name} resources={resources} />
        </Box>
    </div>
}