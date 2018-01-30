import React from 'react'
import {Bold, Button, Caption, Container, ContainerViewContent, Grid, GridItem, H3, Sub1, Sub2, Spacer} from 'hig-react'


import {VerticalTimeline, VerticalTimelineElement} from 'react-vertical-timeline-component';
import 'react-vertical-timeline-component/style.min.css';
import SceneCreationSection from './SceneCreationSection'
import ImageUploadingSection from './ImageUploadingSection';
import SceneStartSection from './SceneStartSection'
import ResultsSection from './ResultsSection'


const header =
    <div style={{margin: "auto", width: "800px"}}>
        <Container >

        <H3>The Reality Capture API provides the photogrammetry capability to process digital images</H3>
            <div style={{textAlign:"center"}}><Sub1>into high resolution textured meshes, dense point clouds and orthophotos.</Sub1></div>

        <Sub1> This sample illustrates how any REST-capable mobile, desktop or web application can connect to this API.</Sub1>

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
    <div>
        <Sub1>Results are available as a time-limited HTTPS link to an output file of the specified format</Sub1>
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
            available_formats: []
        }

    }

    submitSceneCreation = (sceneSettings) => {
        this.setState({
            image_upload: true,
            available_formats: sceneSettings.output_formats
        });
        this.refs["image_uploading_section"].scrollIntoView({behavior: 'smooth'});

    };

    uploadImages = (imageList) => {
        this.setState({
            start_processing: true
        });
        this.refs["start_processing_section"].scrollIntoView({behavior: 'smooth'});
    };


    checkProgress = () => {
        let progress = setInterval(() => {

            //TODO: fetch progress and update the state with it
            this.setState({
                process_progress: this.state.process_progress + 5
            });

            if (this.state.process_progress > 99) {
                clearInterval(progress);
                this.setState({
                    get_results: true
                });
                this.refs["get_results_section"].scrollIntoView({behavior: 'smooth'});

            }
        }, 10);


    };


    requestResultOfType = (format) => {
        console.log("Was requested result of type: " + format);
    };




    startSceneProcessing = () => {
        this.checkProgress();
    };

    render() {
        return (

            <div>

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
                            date="Download results"
                            iconStyle={{background: '#0696D7', color: '#666666'}}
                            // icon={<WorkIcon />}
                        >
                            <ResultsSection formats={this.state.available_formats} getResult={this.requestResultOfType}/>
                    </VerticalTimelineElement>
                    </div>

                </VerticalTimeline>
                <Container>
                    <H3>Finish</H3>
                </Container>
            </div>
        )
    }
}

export default Content;