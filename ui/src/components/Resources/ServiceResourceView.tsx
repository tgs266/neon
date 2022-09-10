import { Box } from "@mui/system";
import React, { useEffect, useState } from "react";
import TitleCard from "../DetailsTitleCard";
import { V1Service } from "@kubernetes/client-node"
import { KubernetesService } from "../../services/Kubernetes";
import Statistic from "../Statistic";
import { Chip, Table, TableBody, TableCell, TableHead, TableRow, Tooltip, Typography } from '@mui/material';
import PodContainerView from "./PodContainerView";



export function ServiceResourceView(props: { namespace: string, name: string }) {

    const [service, setService] = useState<V1Service>(null)

    useEffect(() => {
        KubernetesService.getService(props.namespace, props.name).then(r => {
            setService(r.data)
        })
    }, [props])


    const { name, namespace } = props
    return <Box sx={{ flexGrow: 1, overflow: "auto" }}>
        <TitleCard sx={{mt: 0}} title={name}>
            <Box sx={{ mt: 1, display: "flex", alignItems: "flex-start", gap: 2, flexWrap: "wrap" }}>
                <Statistic label="Created At" value={new Date(service?.metadata.creationTimestamp).toLocaleString()} />
                <Statistic label="Namespace" value={service?.metadata?.namespace} />
                <Statistic label="UID" value={service?.metadata?.uid} />
            </Box>
            <Box sx={{ mt: 1 }}>
                {service?.metadata?.labels && <Statistic label={"Labels"} value={
                    <div style={{ margin: "-4px" }}>
                        {Object.keys(service?.metadata.labels).map(k =>
                            <Chip size="small" label={`${k}: ${service?.metadata.labels[k]}`} sx={{ m: 0.5 }} />
                        )}
                    </div>
                } />}
            </Box>
            <Box sx={{ mt: 1 }}>
                {service?.metadata?.annotations && <Statistic label={"Annotations"} value={
                    <div style={{ margin: "-4px" }}>
                        {Object.keys(service?.metadata.annotations).map(k =>
                            <Chip size="small" label={`${k}: ${service?.metadata.annotations[k]}`} sx={{ m: 0.5 }} />
                        )}
                    </div>
                } />}
            </Box>
        </TitleCard>
        {/* <TitleCard sx={{ mb: 1 }} title="Conditions">
            <div style={{ border: "1px solid rgba(224, 224, 224, 1)", borderRadius: "3px"}}>
            <Table sx={{ minWidth: 650 }}>
                <TableHead>
                    <TableRow>
                        <TableCell>Type</TableCell>
                        <TableCell align="right">Status</TableCell>
                        <TableCell align="right">Reason</TableCell>
                        <TableCell align="right">Message</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {pod?.status?.conditions && pod?.status?.conditions.map((row) => (
                        <TableRow
                            key={row.type}
                            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                            <TableCell component="th" scope="row">
                                {row.type}
                            </TableCell>
                            <TableCell align="right" component="th" scope="row">
                                {row.status}
                            </TableCell>
                            <TableCell align="right" component="th" scope="row">
                                {row.reason}
                            </TableCell>
                            <TableCell align="right" component="th" scope="row">
                                {row.message}
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
            </div>
        </TitleCard>
        <Typography variant="h6" color="text.secondary" sx={{mb: 1, ml: 2}}>Containers</Typography>
        {pod?.spec?.containers.map((_, i) => <PodContainerView idx={i} pod={pod} />)} */}
    </Box>
}