## VisionTrack

### Overview

**VisionTrack** is an advanced computer vision project focused on real-time person detection, tracking, and counting, showcased in an interactive **Streamlit** web app. This project leverages **YOLO (You Only Look Once)** for person detection and integrates the **supervision** library for tracking and counting. Designed to support multiple video streams simultaneously, the system identifies and tracks people in video feeds, maintaining an accurate count over time and providing an intuitive interface for real-time analysis and visualization.

### Learning Objectives

The primary goal of **VisionTrack** is to develop practical skills in building and deploying a real-time, multi-stream computer vision system. By completing this project, you will:

- Implement person detection using **YOLO**.
- Integrate **supervision** for tracking and counting detected individuals.
- Develop and deploy an interactive **Streamlit** app capable of handling multiple video streams.
- Optimize the app for smooth, real-time performance with **GPU acceleration**.
- Enhance understanding of performance optimization through **transfer learning** and model tuning.

### Instructions

#### Data Loading and Preprocessing

1. **Video/Image Dataset Preparation**:
   - Select or capture video streams or images featuring people.
   - Preprocess the data (e.g., resizing, normalization) to ensure it is ready for model input.

2. **Annotation and Labeling**:
   - Use annotation tools like **LabelImg** or **Roboflow** for labeling, if additional training data is needed.
   - Confirm annotations are compatible with YOLO models.

#### Model Implementation

1. **Person Detection with YOLO**:
   - Use a **pre-trained YOLO model** (e.g., YOLOv5 or YOLOv8) tailored for person detection.
   - Fine-tune the model using **transfer learning** to adapt to specific dataset variations with **PyTorch** or **TensorFlow/Keras**.

2. **Integration with Supervision Library**:
   - Apply the **supervision** library for real-time tracking of detected individuals.
   - Manage unique IDs and state tracking across video frames to maintain individual object continuity.

#### Multi-Stream Object Tracking and Counting

1. **Multi-Stream Tracking Pipeline**:
   - Extend the **supervision** library integration to support multiple video streams, enabling simultaneous tracking and analysis.
   - Develop logic to ensure accurate tracking and distinct object IDs for each stream.

2. **Counting Mechanism**:
   - Implement object counting logic using **supervision** to count individuals entering and exiting designated ROIs (regions of interest).
   - Display counts dynamically on each video feed.

#### Streamlit App Development

1. **App Layout**:
   - Create a **Streamlit** app (`app.py`) to display real-time video feeds with overlaid detection, tracking, and counting for multiple streams.
   - Design an intuitive interface that allows users to:
     - Upload or connect multiple video feeds.
     - Toggle between different video streams for detailed analysis.
     - View detection, tracking, and counting results with overlaid bounding boxes and counts.

2. **Interactive Controls**:
   - Include controls for configuring detection thresholds and switching between video streams.
   - Allow toggling of detection, tracking, and counting features for each stream independently.

3. **Visualization**:
   - Use **OpenCV** and **Streamlit** for video rendering with real-time overlays of tracking IDs and counts.
   - Ensure smooth and synchronized rendering of multiple streams.

#### Performance Optimization

1. **Hardware Utilization**:
   - Implement **GPU acceleration** using **CUDA** with **OpenCV** and **PyTorch** to support real-time multi-stream processing.

2. **Model Optimization**:
   - Apply **model pruning** and **quantization** techniques to enhance inference speed and reduce latency, essential for multi-stream performance.

3. **Streamlit App Efficiency**:
   - Optimize the app to manage high-resolution video inputs with minimal lag across multiple streams.

#### Visualization and Analysis

1. **Real-Time Results**:
   - Display detection, tracking, and counting results across all active video streams with clear overlays.
   - Present real-time metrics such as FPS and latency for each stream within the app interface.

2. **Performance Metrics**:
   - Evaluate the app's performance with multi-stream support using metrics like **precision**, **recall**, and **F1-score**.
   - Display performance analysis within the app to inform users of the detection and tracking accuracy.

#### Validation

To ensure project completeness and audit validation, include the following:

1. **Model Artifacts**:
   - Save all trained and optimized YOLO model weights in:
     ```
     models/checkpoints/
     ├── best.pt
     ├── best_quantized.onnx
     └── config.yaml
     ```
   - Include logs or configuration files documenting training and optimization steps.

2. **Evaluation Metrics**:
   - Generate and save a report file:
     reports/performance_metrics.json
   - Example format:
     ```json
     {
       "detection_precision": 0.92,
       "detection_recall": 0.9,
       "f1_score": 0.91,
       "average_fps_per_stream": 18.5,
       "average_latency_ms": 85.0
     }
     ```
   - Minimum passing thresholds:
     Precision ≥ 0.85
     Recall ≥ 0.80
     F1-score ≥ 0.85
     Average FPS ≥ 15 (for 720p video)

3. **Real-Time App Test**
   - The app must run using:
     ```
     streamlit run app.py
     ```
   - The app should:
     Display real-time detection overlays and FPS/latency counters.
     Allow toggling of detection and tracking features per stream.
     Handle missing or broken video sources gracefully.

4. **ROI Counting Validation**
   - Demonstrate ROI-based counting of people entering/exiting the region.
   - Save examples in:
     ```
     reports/demo_results/
     ├── roi_counting_example.png
     └── multi_stream_demo.mp4
     ```

5. **GPU and Fallback Test**
   - Check for CUDA availability in your code:
     ```
     import torch
     print("Using CUDA:", torch.cuda.is_available())
     ```
   - The app must still run on CPU if CUDA is unavailable (with lower FPS

6. **Error Handling**
   - The app must not crash on missing files or failed streams.

   - Log errors to:
     ```
     logs/app_errors.log
     ```

### Project Repository Structure

```
vision-track/
│
├── data/
│   ├── raw_videos/
│   ├── raw_images/
│   └── coco_dataset/
│
├── models/
│   ├── yolo_person_detection.py
│   └── __init__.py
│   └── /checkpoints/
│     ├── best.pt
│     ├── best_quantized.onnx
│     └── config.yaml
│
├── utils/
│   ├── data_loader.py
│   ├── preprocessing.py
│   ├── multi_stream_tracking_helpers.py
│   ├── counting_logic.py
│   ├── VisionTrack_Analysis.ipynb
│   └── __init__.py
│
├── reports/demo_results/
│    ├── roi_counting_example.png
│    └── multi_stream_demo.mp4
│
├── logs/app_errors.log
├── app.py
├── README.md                # Project overview and setup instructions
└── requirements.txt         # List of dependencies
```

### Tips

1. **Pre-Trained Model Advantage**:
   - Start with a pre-trained YOLO model to save time and ensure strong baseline performance for person detection.

2. **Optimize for Multi-Stream Processing**:
   - Ensure the app handles multiple video feeds efficiently by testing on different video sources and resolutions.

3. **User Experience**:
   - Design the app to make switching between streams and accessing real-time analysis seamless and user-friendly.

### Resources

- [YOLOv5 Documentation](https://github.com/ultralytics/yolov5)
- [Supervision Library Documentation](https://github.com/roboflow/supervision)
- [Streamlit Documentation](https://docs.streamlit.io/)
- [PyTorch Documentation](https://pytorch.org/docs/stable/index.html)
- [OpenCV Documentation](https://docs.opencv.org/)
- [CUDA Toolkit](https://developer.nvidia.com/cuda-toolkit)
