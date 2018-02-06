import React from 'react'
import {Bold, Caption, Container, H3, Sub1, Sub2} from 'hig-react'


import {VerticalTimeline, VerticalTimelineElement} from 'react-vertical-timeline-component';
import 'react-vertical-timeline-component/style.min.css';
import SceneCreationSection from './SceneCreationSection'
import ImageUploadingSection from './ImageUploadingSection';
import SceneStartSection from './SceneStartSection';
import ResultsSection from './ResultsSection';


const header =
    <div style={{margin: "auto", width: "800px"}}>
        <Container>

            <H3>The Reality Capture API provides the photogrammetry capability to process digital images</H3>
            <div style={{textAlign: "center"}}><Sub1>into high resolution textured meshes, dense point clouds and
                orthophotos.</Sub1></div>

            <Sub1> This sample illustrates how any REST-capable mobile, desktop or web application can connect to this
                API.</Sub1>

        </Container>
    </div>;

const footer =
    <div style={{margin: "auto", width: "800px"}}>
        <Container>

            <div style={{textAlign: "center"}}>
                <Sub1> This concludes the process and now the photoscene can be deleted.</Sub1>
            </div>

            <Sub2>The resulting Reality Capture data (RCM, OBJ, RCS, GeoTIFF) can be viewed
                within various Autodesk desktop applications such as: ReCap Photo, ReCap Pro, Civil 3D and InfraWorks,
                but also integrated into Forge pipeline and visualized using Forge Viewer.</Sub2>
        </Container>
    </div>;

const scene_creation_comment =
    <div>
        <Sub1>Creates and initializes a photoscene for reconstruction.</Sub1>
        <Sub2>A “photoscene” is a container for a photo-to-3D project.</Sub2>
    </div>;

const image_upload_comment =
    <div>
        <Sub1>We have to add one or more files to a photoscene.</Sub1>
        <Sub2>Files can be added to photoscene either by uploading them </Sub2>
        <Sub2>directly or by providing public HTTP/HTTPS links. </Sub2>

        <div style={{marginLeft: "20%"}}>
            <div style={{textAlign: "center", marginBottom: "0px"}}><Bold>Limitations:</Bold></div>
            <ul style={{textAlign: "left", marginTop: "0px", marginLeft: "10%"}}>
                <li><Caption>only JPEG images are supported.</Caption></li>
                <li><Caption>maximum number of files in a single request: 20</Caption></li>
                <li><Caption>maximum size of a single file: 128 MB</Caption></li>
                <li><Caption>maximum uncompressed size of image in memory: 512 MB</Caption></li>
            </ul>
        </div>

    </div>;


//TODO: Change this comment to something more useful
const scene_progress_comment =
    <div>
        <Sub1>Check status by calling: GET photoscene/:photosceneid/progress</Sub1>

        <Sub2>Returns the processing progress and status of a photoscene.</Sub2>
    </div>;


const scene_download_comment =
    <div style={{marginLeft: "20%"}}>

        <Sub1>Download results</Sub1>
        <Sub2>Results are available as a time-limited HTTPS link to an output file</Sub2>
        <Sub2>of the specified format.</Sub2>
        <Sub2>The link will expire 7 days after the date of processing completion.</Sub2>
    </div>;

class Content extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            scene_setup: true,
            image_upload: false,
            start_processing: false,
            process_progress: 0,
            get_results: false,
            finished: false,
            available_formats: [],
            scene_id: ""
        };
    }

    componentDidMount() {
        this.refs["header"].scrollIntoView({behavior: 'smooth'});
    }

    submitSceneCreation = (sceneSettings) => {
        this.setState({
            available_formats: sceneSettings.output_formats
        });

        fetch("http://" + window.location.hostname + ":3000/create_scene", {
            method: 'POST',
            body: JSON.stringify(sceneSettings),
        }).then(res => res.json()).then(response => {
            console.log("Scene settings sent successfully, got: ", JSON.stringify(response));
            this.setState({
                image_upload: true,
                scene_id: response["scene_id"]
            });
            this.refs["bottom"].scrollIntoView({behavior: 'smooth'});

        }).catch(error => {
            console.log("Could not send sceneSettings: ", error)
        });
    };

    uploadImages = (imageList) => {

        fetch("http://" + window.location.hostname + ":3000/upload_remote_images", {
            method: 'POST',
            body: JSON.stringify({
                scene_id: this.state.scene_id,
                image_list: imageList
            }),
        }).then(res => res.json()).then(response => {

            console.log("Image list successfully, got: ", JSON.stringify(response));
            this.setState({
                start_processing: true
            });
            this.refs["bottom"].scrollIntoView({behavior: 'smooth'});

        }).catch(error => {
            console.log("Could not send list of remote images: ", error)
        });
    };


    checkProgress = () => {
        let progress = setInterval(() => {
            fetch("http://" + window.location.hostname + ":3000/check_progress", {
                method: 'POST',
                body: JSON.stringify({
                    scene_id: this.state.scene_id
                }),
            }).then(res => res.json()).then(response => {
                console.log("Scene progress check: ", response);
                this.setState({
                    process_progress: response["progress"]
                });
            }).catch(error => {
                console.log("Could not start scene processing: ", error)
            });


            if (this.state.process_progress > 99) {
                clearInterval(progress);
                this.setState({
                    get_results: true,
                    finished: true
                });
                this.refs["bottom"].scrollIntoView({behavior: 'smooth'});

            }
        }, 1000);


    };


    requestResultOfType = (format) => {
        console.log("Was requested result of type: " + format);
        fetch("http://" + window.location.hostname + ":3000/get_result", {
            method: 'POST',
            body: JSON.stringify({
                scene_id: this.state.scene_id,
                format: format
            }),
        }).then(res => res.json()).then(response => {

            console.log("Received scene result: ", JSON.stringify(response));

        }).catch(error => {
            console.log("Could not send list of remote images: ", error)
        });
    };


    startSceneProcessing = () => {
        fetch("http://" + window.location.hostname + ":3000/start_process", {
            method: 'POST',
            body: JSON.stringify({
                scene_id: this.state.scene_id
            }),
        }).then(res => res.json())
            .catch(error => {
                console.log("Could not start scene processing: ", error)
            })
            .then(response => {
                if (response !== undefined && response["result"] === "ACK") {
                    console.log("Scene processing started successfully: ", JSON.stringify(response));
                    this.setState({
                        start_processing: true
                    });
                    this.refs["bottom"].scrollIntoView({behavior: 'smooth'});
                } else {
                    console.log("Could not properly start scene processing")
                }

            });
        this.checkProgress();
    };

    render() {

        let show_footer = this.state.finished ? footer : "";


        return (

            <div ref="header">

                {header}


                <VerticalTimeline>
                    <VerticalTimelineElement
                        className="scene_creation_element"
                        date={scene_creation_comment}
                        iconStyle={{background: '#0696D7', color: '#fff'}}
                        // icon={<WorkIcon />}
                    >
                        <SceneCreationSection submit={this.submitSceneCreation}/>
                    </VerticalTimelineElement>

                    <div ref="image_uploading_section" style={{display: this.state.image_upload ? "inline" : "none"}}>
                        <VerticalTimelineElement
                            className="image_uploading_element"
                            date={image_upload_comment}
                            iconStyle={{background: '#0696D7', color: '#666666'}}
                            style={{display: this.state.image_upload ? "inline" : "none"}}
                            // icon={<WorkIcon />}
                            position="right"
                        >

                            <ImageUploadingSection upload={this.uploadImages}/>

                        </VerticalTimelineElement>
                    </div>

                    <div ref="start_processing_section"
                         style={{display: this.state.start_processing ? "inline" : "none"}}>
                        <VerticalTimelineElement
                            className="start_processing_element"
                            date={scene_progress_comment}
                            iconStyle={{background: '#0696D7', color: '#666666'}}
                            // icon={<WorkIcon />}
                        >

                            <SceneStartSection progress={this.state.process_progress}
                                               startProcess={this.startSceneProcessing}/>


                        </VerticalTimelineElement>
                    </div>

                    <div ref="get_results_section" style={{display: this.state.get_results ? "inline" : "none"}}>
                        <VerticalTimelineElement
                            position="right"
                            className="query_results_element"
                            date={scene_download_comment}
                            iconStyle={{background: '#0696D7', color: '#666666'}}
                            // icon={<WorkIcon />}
                        >
                            <ResultsSection formats={this.state.available_formats}
                                            getResult={this.requestResultOfType}/>
                        </VerticalTimelineElement>
                    </div>

                </VerticalTimeline>

                {show_footer}

                <Container>
                    <div ref="bottom">
                        <Caption>&copy; Autodesk Forge 2018</Caption>
                    </div>
                </Container>
            </div>
        )
    }
}

export default Content;