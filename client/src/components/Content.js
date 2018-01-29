import React from 'react'
import {Button, Container, ContainerViewContent, Grid, GridItem, H3} from 'hig-react'

import {VerticalTimeline, VerticalTimelineElement} from 'react-vertical-timeline-component';
import 'react-vertical-timeline-component/style.min.css';
import SceneCreationSection from './SceneCreationSection'
import ImageUploadingSection from './ImageUploadingSection';
import SceneStartSection from './SceneStartSection'
import ResultsSection from './ResultsSection'

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
        console.table(imageList);
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


    startSceneProcessing = () => {
        this.checkProgress();
        console.table(this.state)
    };

    render() {
        return (

            <Container>

                <Container>
                    Setup
                </Container>


                <VerticalTimeline>
                    <VerticalTimelineElement
                        className="scene_creation_element"
                        date="Create a scene"
                        iconStyle={{background: '#0696D7', color: '#fff'}}
                        // icon={<WorkIcon />}
                    >
                        <SceneCreationSection submit={this.submitSceneCreation}/>
                    </VerticalTimelineElement>

                    <div ref="image_uploading_section" style={{display: this.state.image_upload ? "inline" : "none"}}>
                        <VerticalTimelineElement
                            className="image_uploading_element"
                            date="Resource uploading"
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
                            date="Initiate work"
                            iconStyle={{background: '#0696D7', color: '#666666'}}
                            // icon={<WorkIcon />}
                        >

                            <SceneStartSection progress={this.state.process_progress} startProcess={this.startSceneProcessing}/>


                        </VerticalTimelineElement>
                    </div>

                    <div ref="get_results_section" style={{display: this.state.get_results ? "inline" : "none"}}>
                        <VerticalTimelineElement
                            position="right"
                            className="query_results_element"
                            date="Processing"
                            iconStyle={{background: '#0696D7', color: '#666666'}}
                            // icon={<WorkIcon />}
                        >
                            <ResultsSection formats={this.state.available_formats}/>
                        </VerticalTimelineElement>
                    </div>

                </VerticalTimeline>
                <Container>
                    <H3>Finish</H3>
                </Container>
            </Container>
        )
    }
}

export default Content;