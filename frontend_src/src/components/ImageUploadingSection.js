import React from 'react';
import {Button, Container, H3, IconButton, Spacer, Tab, Table, Tabs, TextCellContent, TextField} from 'hig-react';

import DropzoneComponent from 'react-dropzone-component'


const sample_images = [
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5264.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5263.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5262.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5261.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5260.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5259.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5258.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5257.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5256.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5255.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5254.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5253.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5252.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5251.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5250.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5249.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5248.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5247.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5246.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5245.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5244.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5243.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5242.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5241.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5240.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5239.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5238.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5237.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5236.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5235.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5234.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5233.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5232.JPG",
    "https://s3.amazonaws.com/photo-ii-3d-sample-images/DSC_5231.JPG",



];


let dropZoneConfig = {
    iconFiletypes: ['.jpg'],
    showFiletypeIcon: true,
    postUrl: "http://" + window.location.hostname + ":3000/uploadLocalImages"
};


class ImageUploadingSection extends React.Component {
    constructor(props) {
        super(props);

        let sample = [];
        // populate example with sample images
        for (let i = 0; i < sample_images.length; ++i) {
            sample.push({
                "id": i,
                "image_uri": sample_images[i]
            })
        }

        this.state = {
            uri_collection: sample,
            done_uploading: false
        }
    }


    removeUriFromList = (event, id) => {
        // event.preventDefault();
        const {uri_collection} = this.state;
        let index = -1;
        for (let i = 0; i < uri_collection.length; ++i) {
            if (id === uri_collection[i]["id"]) {
                index = i;
                break;
            }
        }


        if (index >= 0) {
            this.setState({
                uri_collection: uri_collection.slice(0, index).concat(uri_collection.slice(index + 1))
            });
        }
    };


    addNewURI = (value) => {
        const {uri_collection} = this.state;

        const index = uri_collection.length;
        this.setState({
            uri_collection: uri_collection.concat({
                "id": index,
                "image_uri": value
            })
        })
    };


    processURI = (event) => {

        //TODO: add here check if it is valid url
        this.addNewURI(event.target.value);
    };


    uploadImages = () => {
        this.setState({
            done_uploading: true
        });

        this.props.upload(this.state.uri_collection);
    };


    doneAddingLocalImages = () => {
        this.setState({
            done_uploading: true
        });
        this.props.upload([]);
    };


    render() {

        const {done_uploading} = this.state;

        let visibility = done_uploading ? "none" : "inline";

        const column_headers_remote = [
            {
                id: "1",
                HeaderCell: "URI",
                alignment: "left",
                width: "1fr",
                accessor: "image_uri",
                Cell: props => (
                    <TextCellContent text={props.data.image_uri} detail={props.data.image_uri}/>

                )
            },
            {
                id: "2",
                HeaderCell: "",
                alignment: "right",
                width: "0.15fr",
                accessor: "delete",
                Cell: props => (
                    <IconButton title="delete"
                                icon="trash"
                                size="small"
                                onClick={(e) => this.removeUriFromList(e, props.data.id)}/>
                )
            }

        ];


        return (
            <div>
                <H3>Upload images</H3>
                {/*<Container>*/}
                <Tabs
                    // activeTabIndex={this.state.activeTabIndex}
                    // onTabChange={this.setActiveTabIndex}
                >
                    <Tab label="from a Remote Server">
                        <div style={{display: visibility}}>
                            <Table
                                // selectable
                                density="compressed"
                                columns={column_headers_remote}
                                data={this.state.uri_collection}

                            />

                            <TextField
                                label="add new image uri"
                                // placeholder=""
                                // instructions="add uri and press Enter"
                                onChange={this.processURI}
                            />
                        </div>
                        <Spacer inset="xxs"/>
                        <div style={{textAlign: "center"}}>
                            <Button title="Upload images" onClick={this.uploadImages} disabled={done_uploading}/>
                        </div>
                    </Tab>
                    <Tab label="from a Local Drive">
                        <Container>
                            <div style={{display: visibility}}>
                                <DropzoneComponent
                                    config={dropZoneConfig}
                                    // eventHandlers={(file) => this.prepareImageForUploading(file)}
                                    // djsConfig={{autoProcessQueue:false}}
                                />
                            </div>
                            <div style={{textAlign: "center"}}>
                                <Spacer inset="xxs"/>
                                <Button title="Done" onClick={this.doneAddingLocalImages} disabled={done_uploading}/>

                            </div>
                        </Container>
                    </Tab>
                </Tabs>


            </div>
        )
    }
}

export default ImageUploadingSection;