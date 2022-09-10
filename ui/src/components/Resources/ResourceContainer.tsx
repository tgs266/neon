import { Card, List, ListItemButton, ListItem, ListItemText } from "@mui/material"
import { borderColor } from "@mui/system";
import React, { useState } from "react"
import { ResourceList } from "../../models/resources"
import ExpandableList from "../List/ExpandableList"
import { PodResourceView } from "./PodResourceView";
import { ServiceResourceView } from "./ServiceResourceView";


interface SelectedResource {
    name: string,
    type: "POD" | "SERVICE"
}


export default function ResourceContainer(props: { resources?: ResourceList, appName: string }) {
    const { resources } = props;

    const [selectedResource, setSelectedResource] = useState<SelectedResource>(null)

    return <div style={{ display: "flex", height: "100%" }}>
        <div style={{
            width: "250px",
            maxHeight: "100%",
            flexShrink: 0,
        }}>
            <Card sx={{height: "calc(100% - 16px)", overflow: "auto", p: 0, mt: 1}}>
                <List disablePadding>
                    <ExpandableList primary="Pods">
                        {resources?.pods && resources?.pods.map(pod =>
                            <ListItemButton onClick={() => setSelectedResource({ name: pod, type: "POD" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={pod} />
                            </ListItemButton>
                        )}
                        {resources?.pods && resources?.pods.map(pod =>
                            <ListItemButton onClick={() => setSelectedResource({ name: pod, type: "POD" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={pod} />
                            </ListItemButton>
                        )}
                        {resources?.pods && resources?.pods.map(pod =>
                            <ListItemButton onClick={() => setSelectedResource({ name: pod, type: "POD" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={pod} />
                            </ListItemButton>
                        )}
                        {resources?.pods && resources?.pods.map(pod =>
                            <ListItemButton onClick={() => setSelectedResource({ name: pod, type: "POD" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={pod} />
                            </ListItemButton>
                        )}
                        {resources?.pods && resources?.pods.map(pod =>
                            <ListItemButton onClick={() => setSelectedResource({ name: pod, type: "POD" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={pod} />
                            </ListItemButton>
                        )}
                        {resources?.pods && resources?.pods.map(pod =>
                            <ListItemButton onClick={() => setSelectedResource({ name: pod, type: "POD" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={pod} />
                            </ListItemButton>
                        )}
                        {resources?.pods && resources?.pods.map(pod =>
                            <ListItemButton onClick={() => setSelectedResource({ name: pod, type: "POD" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={pod} />
                            </ListItemButton>
                        )}
                        {resources?.pods && resources?.pods.map(pod =>
                            <ListItemButton onClick={() => setSelectedResource({ name: pod, type: "POD" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={pod} />
                            </ListItemButton>
                        )}
                    </ExpandableList>
                    <ExpandableList primary="Services">
                        {resources?.services && resources?.services.map(svc =>
                            <ListItemButton onClick={() => setSelectedResource({ name: svc, type: "SERVICE" })} sx={{ pl: 4 }}>
                                <ListItemText primaryTypographyProps={{ fontSize: "14px" }} primary={svc} />
                            </ListItemButton>
                        )}
                    </ExpandableList>
                </List>
            </Card>
        </div>
        {selectedResource?.type === "POD" && <PodResourceView namespace={props.appName} podName={selectedResource?.name} />}
        {selectedResource?.type === "SERVICE" && <ServiceResourceView namespace={props.appName} name={selectedResource?.name} />}
    </div>


}