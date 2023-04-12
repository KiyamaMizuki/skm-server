import React from "react";
// Mapbox GL JS読み込み
// import mapboxgl from "mapbox-gl";
import "mapbox-gl/dist/mapbox-gl.css";
import "./MapPane.css";
import marker_green from "./marker/green.png";
import marker_blue from "./marker/blue.jpg";
import marker_red from "./marker/red.jpg";
import "leaflet/dist/leaflet.css";
import axios from "axios";
import L, { LatLng, Marker } from "leaflet";
import { Modal } from "antd";
import "antd/dist/antd.css";

interface MapState {
  nodeList: [
    {
      ID: number;
      Latitude: number;
      Longitude: number;
      Floor: number;
      Name: string;
      Type: number;
    }
  ];
  lineList: any;
  visible: boolean;
  nodeVisible: boolean;
  name: string;
  type: number;
  lat: number;
  lng: number;
  zoom: number;
  zoomControl: boolean;
  currentFloor: number;
}

export default class MapPane extends React.Component<{}, MapState> {
  constructor(props) {
    super(props);
    this.onChange = this.onChange.bind(this);
    this.changeFloorMinus = this.changeFloorMinus.bind(this);
    this.changeFloorPlus = this.changeFloorPlus.bind(this);
    this.state = {
      nodeList: [
        {
          ID: 0,
          Latitude: 0,
          Longitude: 0,
          Floor: 0,
          Name: "",
          Type: 1,
        },
      ],
      lineList: [
        {
          startLatitude: 0,
          startLongitude: 0,
          endLatitude: 0,
          endLongitude: 0,
        },
      ],
      visible: false,
      nodeVisible: false,
      name: "",
      lat: 26.2517,
      lng: 127.7684,
      type: 1,
      zoom: 17,
      zoomControl: true,
      currentFloor: 1,
    };
  }
  marker = Array();
  map: any;
  container: any;
  m_streets: any;
  modalIsOpen: Boolean | undefined;
  line_list = Array();
  node!:
    | {
        ID: number;
        Latitude: number;
        Longitude: number;
        Floor: number;
        Type: number;
        Name: string;
      }
    | undefined;
  nodeTo!:
    | {
        ID: number;
        Latitude: number;
        Longitude: number;
        Floor: number;
        Type: number;
        Name: string;
      }
    | undefined;

  blueIcon = L.icon({
    iconUrl: marker_blue,
    iconSize: [10, 10],
    iconAnchor: [5, 5],
  });
  redIcon = L.icon({
    iconUrl: marker_red,
    iconSize: [10, 10],
    iconAnchor: [5, 5],
  });
  greenIcon = L.icon({
    iconUrl: marker_green,
    iconSize: [10, 10],
    iconAnchor: [5, 5],
  });
  async getNode() {
    await axios.get("/node").then((results) => {
      const nodeList = results.data;
      this.setState({ nodeList });
      this.state.nodeList.map((data) => {
        if (data.Floor === this.state.currentFloor) {
          switch (data.Type) {
            case 1:
              var marker = new L.Marker(
                [Number(data.Latitude), Number(data.Longitude)],
                {
                  icon: this.blueIcon,
                }
              )
                .on("click", () => this.showModal(data))
                .addTo(this.map);
              this.marker.push(marker);
              break;
            case 2:
              var marker = new L.Marker(
                [Number(data.Latitude), Number(data.Longitude)],
                {
                  icon: this.greenIcon,
                }
              )
                .on("click", () => this.showModal(data))
                .addTo(this.map);
              this.marker.push(marker);
              break;
            case 3:
              var marker = new L.Marker(
                [Number(data.Latitude), Number(data.Longitude)],
                {
                  icon: this.redIcon,
                }
              )
                .on("click", () => this.showModal(data))
                .addTo(this.map);
              this.marker.push(marker);
              break;
          }
        }
      });
    });
  }
  async getRoad() {
    await axios.get("/road").then((results) => {
      console.log(results.data["road"]);
      const roadList = results.data["road"];
      const lineList = new Array();
      roadList.forEach((element) => {
        const startID = element.NodeStartID;
        const endID = element.NodeEndID;
        const floor = element.Floor;
        const startNode = this.state.nodeList.find((v) => v.ID === startID);
        const endNode = this.state.nodeList.find((v) => v.ID === endID);
        // console.log("start-->",startNode,"end-->",endNode);
        const line = {
          startLatitude: startNode?.Latitude,
          startLongitude: startNode?.Longitude,
          endLatitude: endNode?.Latitude,
          endLongitude: endNode?.Longitude,
          floor: floor,
        };
        lineList.push(line);
      });
      // console.log("lineList-->",lineList);
      this.setState({ lineList });
      this.state.lineList.map((data) => {
        if (data.floor === this.state.currentFloor) {
         var line = new L.Polyline([
            [data.startLatitude, data.startLongitude],
            [data.endLatitude, data.endLongitude],
          ]).addTo(this.map);
          this.line_list.push(line)
        }
      });
    });
  }
  async postNode(body) {
    await axios
      .post("/node", body)
      .then((response) => {
        this.node = response.data["node"];
        this.state.nodeList.push(this.node!);
      })
      .catch((error) => console.log(error));
  }
  async componentDidMount() {
    axios.defaults.baseURL = "http://localhost:1323";
    axios.defaults.headers.post["Access-Control-Allow-Origin"] =
      "http://localhost:1323";
    axios.defaults.headers = {
      "Content-Type": "application/x-www-form-urlencoded",
    };
    this.m_streets = L.tileLayer(
      "https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token={accessToken}",
      {
        attribution:
          '© <a href="https://www.mapbox.com/about/maps/">Mapbox</a> © <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a> <strong><a href="https://www.mapbox.com/map-feedback/" target="_blank">Improve this map</a></strong>',
        tileSize: 512,
        maxZoom: 24,
        zoomOffset: -1,
        id: "mapbox/streets-v11",
        accessToken:
          "pk.eyJ1Ijoic2hpbnlhLXRhbiIsImEiOiJja2dhaHJzcTUwN2FqMnlvOWV4MzR4bnBwIn0.IXRxzUcZW3yK9qXRFjUWSQ",
      }
    );
    const body = new URLSearchParams();
    this.map = L.map(this.container, {
      center: [26.2517, 127.7684],
      zoom: 17,
      zoomControl: true,
      layers: [this.m_streets],
    });
    var gson =
      '{"type":"FeatureCollection","features":[{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#555555","fill-opacity":0.5,"name":"階段"},"geometry":{"type":"Polygon","coordinates":[[[127.76668142527343,26.253435080541127],[127.76672065258026,26.253448310950024],[127.76670154184103,26.253489806313564],[127.76666432619095,26.253477477982464],[127.76668142527343,26.253435080541127]]]}},{"type":"Feature","properties":{},"geometry":{"type":"LineString","coordinates":[[127.76666767895222,26.253471764852975],[127.76670422405005,26.25348319111167]]}},{"type":"Feature","properties":{},"geometry":{"type":"LineString","coordinates":[[127.76667036116123,26.253464848959002],[127.76670657098295,26.25347597452737]]}},{"type":"Feature","properties":{},"geometry":{"type":"LineString","coordinates":[[127.76667371392249,26.25345612891821],[127.76671059429646,26.25346815656051]]}},{"type":"Feature","properties":{},"geometry":{"type":"LineString","coordinates":[[127.76668645441534,26.253476575909396],[127.76670154184103,26.253440793672446]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#555555","fill-opacity":0.5,"name":"工1-121"},"geometry":{"type":"Polygon","coordinates":[[[127.76647992432116,26.25335810358684],[127.76657279580832,26.253393885849288],[127.76653356850147,26.25347807936449],[127.76643969118594,26.25344410127464],[127.76647992432116,26.25335810358684]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#555555","fill-opacity":0.5,"name":"工1-122"},"geometry":{"type":"Polygon","coordinates":[[[127.76657212525608,26.25339478792299],[127.76664454489945,26.253420045983756],[127.7666049823165,26.253503638098127],[127.76653390377759,26.25347807936449],[127.76657212525608,26.25339478792299]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#fcfcfc","fill-opacity":0.5},"geometry":{"type":"Polygon","coordinates":[[[127.76653725653887,26.253312699187532],[127.76666834950446,26.25336411741322],[127.76664689183235,26.253407416954058],[127.76651713997126,26.25335810358684],[127.76653725653887,26.253312699187532]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#555555","fill-opacity":0.5,"name":"工学部事務"},"geometry":{"type":"Polygon","coordinates":[[[127.76627540588379,26.25315212984499],[127.76650540530683,26.2532384283769],[127.7664802595973,26.253290448018483],[127.76625026017427,26.253203247450045],[127.76627540588379,26.25315212984499]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#555555","fill-opacity":0.5,"name":"階段"},"geometry":{"type":"Polygon","coordinates":[[[127.7661268785596,26.25319693292357],[127.76617180556057,26.253212268201544],[127.7661657705903,26.25322429586911],[127.76616409420967,26.25322970831911],[127.76616074144839,26.253235722152155],[127.76611547917128,26.253220086185607],[127.7661268785596,26.25319693292357]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#555555","fill-opacity":0.5,"name":"工1-111"},"geometry":{"type":"Polygon","coordinates":[[[127.76615470647812,26.25325797333169],[127.76630155742167,26.25331209780464],[127.76627004146576,26.25338005405167],[127.76612386107445,26.25332653099319],[127.76615470647812,26.25325797333169]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#555555","fill-opacity":0.5,"name":"工1-112"},"geometry":{"type":"Polygon","coordinates":[[[127.76630155742167,26.25331209780464],[127.76643969118594,26.253364718795858],[127.76640951633452,26.253431171556397],[127.76627004146576,26.25338005405167],[127.76630155742167,26.25331209780464]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#5d62ee","fill-opacity":0.5,"name":"男子トイレ"},"geometry":{"type":"Polygon","coordinates":[[[127.76664454489945,26.253419143910257],[127.76666164398193,26.2534269618804],[127.76664856821299,26.253456429609273],[127.76663113385439,26.253448912332203],[127.76664454489945,26.253419143910257]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#ee4949","fill-opacity":0.5,"name":"女子トイレ"},"geometry":{"type":"Polygon","coordinates":[[[127.76666197925805,26.253425458424644],[127.76668108999728,26.253434178467753],[127.76666902005672,26.253464247576897],[127.76664823293684,26.25345733168249],[127.76666197925805,26.253425458424644]]]}},{"type":"Feature","properties":{"stroke":"#555555","stroke-width":2,"stroke-opacity":1,"fill":"#ffffff","fill-opacity":0.5},"geometry":{"type":"Polygon","coordinates":[[[127.7662492543459,26.25322219102736],[127.7664527669549,26.2532961611569],[127.76643365621567,26.253339761414427],[127.76622846722601,26.25326428785483],[127.7662492543459,26.25322219102736]]]}}]}';
    var geo2 = JSON.parse(gson);
    var geoL2 = L.geoJSON(geo2).addTo(this.map);
    await this.getNode();
    await this.getRoad();
    this.map.on("click", async (e: { latlng: { lat: any; lng: any } }) => {
      const lat = e.latlng.lat;
      const lng = e.latlng.lng;
      this.showInput(lat, lng);
      switch (this.state.type) {
        case 1:
          new L.Marker([lat, lng], { icon: this.blueIcon })
            .on("click", () => this.showModal(this.node))
            .addTo(this.map);
          break;
        case 2:
          new L.Marker([lat, lng], { icon: this.blueIcon })
            .on("click", () => this.showModal(this.node))
            .addTo(this.map);
          break;
        case 3:
          new L.Marker([lat, lng], { icon: this.blueIcon })
            .on("click", () => this.showModal(this.node))
            .addTo(this.map);
          break;
      }
    });
  }
  async onItemClick() {}
  showModal(node) {
    this.node = node;
    this.setState({
      visible: true,
    });
  }
  showInput(lat, lng) {
    this.setState({
      nodeVisible: true,
      lat: lat,
      lng: lng,
    });
  }
  onChange(e) {
    this.nodeTo = this.state.nodeList.find(
      (value) => String(value.ID) + " " + value.Name === String(e.target.value)
    );
  }
  async changeFloorPlus(e) {
    const floor = this.state.currentFloor + 1;
    this.marker.forEach(element => {
      this.map.removeLayer(element)
    })
    this.line_list.forEach(element => {
      this.map.removeLayer(element)
    });
    this.marker = new Array()
    this.line_list = new Array()
    this.setState({
      currentFloor: floor,
    });
    await this.getNode();
    await this.getRoad();
  }
  async changeFloorMinus(e) {
    const floor = this.state.currentFloor - 1;
    this.marker.forEach(element => {
      this.map.removeLayer(element)
    })
    this.line_list.forEach(element => {
      this.map.removeLayer(element)
    });
    this.marker = new Array()
    this.line_list = new Array()
    this.setState({
      currentFloor: floor,
    });
    await this.getNode();
    await this.getRoad();
  }
  handleOk = async () => {
    const p1 = new LatLng(
      Number(this.node?.Latitude),
      Number(this.node?.Longitude)
    );
    const p2 = new LatLng(
      Number(this.nodeTo?.Latitude),
      Number(this.nodeTo?.Longitude)
    );
    const distance = p1.distanceTo(p2);
    const body = new URLSearchParams();
    console.log(distance);
    body.append("distance", String(distance));
    body.append("floor", String(this.state.currentFloor));
    body.append("nodeid1", String(this.node?.ID));
    body.append("nodeid2", String(this.nodeTo?.ID));
    await axios.post("/road", body).then(() => {
      new L.Polyline([
        [this.node!.Latitude, this.node!.Longitude],
        [this.nodeTo!.Latitude, this.nodeTo!.Longitude],
      ]).addTo(this.map);
    });
    this.setState({
      visible: false,
    });
  };

  handleCancel = () => {
    this.setState({
      visible: false,
      nodeVisible: false,
    });
  };

  getNodeBody = async () => {
    const lat = this.state.lat;
    const lng = this.state.lng;
    const body = new URLSearchParams();
    body.append("latitude", String(lat));
    body.append("longitude", String(lng));
    body.append("floor", String(this.state.currentFloor));
    body.append("node_type", String(this.state.type));
    body.append("name", this.state.name);
    console.log(this.state.name);
    await this.postNode(body);
    this.setState({
      nodeVisible: false,
    });
  };

  render() {
    return (
      <div className={"mapcontainer"}>
        <div className={"map"} ref={(e) => (this.container = e)}>
          <div>
            <Modal
              title="接続するノードを選択してください"
              visible={this.state.visible}
              onOk={this.handleOk}
              onCancel={this.handleCancel}
            >
              <a>
                【ID】{this.node?.ID} 【名前】{this.node?.Name}
              </a>
              <br></br>
              <select id="dropdown" onChange={this.onChange}>
                {this.state.nodeList.map((option, key) => (
                  <option key={key}>
                    {option.ID} {option.Name}
                  </option>
                ))}
              </select>
            </Modal>
          </div>
          <div>
            <Modal
              title="Nodeの名前とtypeを入力してください"
              visible={this.state.nodeVisible}
              onOk={this.getNodeBody}
              onCancel={this.handleCancel}
            >
              <input
                type="text"
                name="nodeName"
                placeholder=""
                onChange={(e) => this.setState({ name: e.target.value })}
              />
              <br></br>
              <br></br>
              <input
                type="number"
                name="nodeType"
                placeholder=""
                onChange={(e) =>
                  this.setState({ type: Number(e.target.value) })
                }
              />
            </Modal>
          </div>
        </div>

        <p>current floor {this.state.currentFloor}</p>
        <img className={"blueIcon"} src={marker_blue}></img>
        <p>1, 教室</p>
        <img className={"greenIcon"} src={marker_green}></img>
        <p className={"greenIconText"}>2, 道路・玄関</p>
        <img className={"redIcon"} src={marker_red}></img>
        <p className={"redIconText"}>3, 階段・エレベーター</p>

        <button className={"plus_floor"} onClick={this.changeFloorPlus}>
          change floor. +1
        </button>
        <button className={"minus_floor"} onClick={this.changeFloorMinus}>
          change floor. -1
        </button>
      </div>
    );
  }
}
