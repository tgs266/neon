import { Box } from "@mui/system";
import React, { useEffect, useState } from "react";
import TitleCard from "../DetailsTitleCard";
import { V1Pod, V1ContainerState } from "@kubernetes/client-node"
import { KubernetesService } from "../../services/Kubernetes";
import Statistic from "../Statistic";
import { Chip, Table, TableBody, TableCell, TableHead, TableRow, Tooltip, Typography } from '@mui/material';


interface ActualState {
    type: "running" | "waiting" | "terminated"
    startedAt?: Date
    message?: string
    reason?: string
    restartCount?: number
}


function getContainerState(containerState?: V1ContainerState): ActualState {
    if (!containerState) {
        return
    }
    if (containerState.running) {
        return {
            type: "running",
            startedAt: containerState.running.startedAt,
        }
    } else if (containerState.waiting) {
        return {
            type: "waiting",
            message: containerState.waiting.message,
            reason: containerState.waiting.reason
        }
    } else {
        return {
            type: "terminated",
            message: containerState.terminated.message,
            reason: containerState.terminated.reason
        }
    }
}


export default function PodContainerView(props: { pod: V1Pod, idx: number }) {
    const { pod, idx } = props

    const containerSpec = pod?.spec?.containers[idx]
    const containerStatus = pod?.status?.containerStatuses[idx]

    const actualContainerState = getContainerState(containerStatus.state)

    return <TitleCard sx={{ mb: 1 }} title={containerSpec.name}>
        <Box sx={{ mt: 1, display: "flex", alignItems: "flex-start", gap: 2, flexWrap: "wrap" }}>
            <Statistic label="Image" value={containerSpec.image} />
        </Box>
        <Box sx={{ mt: 1, display: "flex", alignItems: "flex-start", gap: 2, flexWrap: "wrap" }}>
            <Statistic label="Ready" value={containerStatus.ready.toString()} />
            <Statistic label="Started" value={containerStatus.started.toString()} />
            {actualContainerState.type === "running" && <Statistic label="Started At" value={actualContainerState.startedAt.toLocaleString()} />}
            <Statistic label="Restart Count" value={containerStatus.restartCount} />
        </Box>
        <Box sx={{ mt: 1, display: "flex", alignItems: "flex-start", gap: 2, flexWrap: "wrap" }}>
            <Statistic label="Reason" value={actualContainerState.reason} />
            <Statistic label="Message" value={actualContainerState.message} />
        </Box>
        <Box sx={{ mt: 1 }}>
            <Statistic label="Mounts" value={
                <div style={{ border: "1px solid rgba(224, 224, 224, 1)", borderRadius: "3px"}}>
                <Table sx={{ width: "100%" }}>
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell align="right">Read Only</TableCell>
                            <TableCell align="right">Mount Path</TableCell>
                            <TableCell align="right">Sub Path</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {containerSpec?.volumeMounts && containerSpec?.volumeMounts.map((row) => (
                            <TableRow
                                key={row?.name}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell component="th" scope="row">
                                    {row?.name}
                                </TableCell>
                                <TableCell align="right" component="th" scope="row">
                                    {row?.readOnly}
                                </TableCell>
                                <TableCell align="right" component="th" scope="row">
                                    {row?.mountPath}
                                </TableCell>
                                <TableCell align="right" component="th" scope="row">
                                    {row?.subPath}
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
                </div>
            } />
        </Box>
    </TitleCard>
}