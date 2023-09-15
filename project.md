## Earth Lens
---

### Pitch

An easy-to-use, beautiful dashboard that retrieves information local to a specific area regarding the environmental landscape.

### Requirements

#### Data Gathering

- Make use of several APIs for various metrics, including the following:
  - Air Quality: Retrieve metrics like PM2.5, PM10, NO2, SO2.
  - Water Quality: Access data on pollutants like heavy metals, microbial contaminants.
  - Land Use: Data on urban sprawl, agricultural practices, and protected areas.
  - Waste Management: Information on recycling rates, landfill waste, and composting.
  - Energy Use: Percentage of energy from renewable sources, energy efficiency metrics.
  - Transportation: Public transport use statistics, EV charging stations, bike lanes.
  - Legislation: Presence and effectiveness of environmental laws.
  - "Green" Careers: Stats on jobs in renewable energy, conservation, and other green sectors.

#### Real-Time Updates

- Make use of Web Sockets or Server Sent Events to provide real-time updates for:
  - Views: Display real-time user interaction metrics.
  - API information: Update environmental metrics every 5 minutes.

#### User Experience

- Geolocation: Allow users to specify their location to get local data.
- User Dashboard: Enable users to customize their dashboard view, selecting which metrics are most relevant to them.
- Mobile Responsiveness: Ensure the dashboard is fully functional on various screen sizes.

#### Alerts and Notifications

- Threshold Alerts: Allow users to set thresholds for specific metrics and send notifications when crossed.
- Legislation Updates: Notify users of any new or updated environmental legislation affecting their area.

#### Data Visualization

- Use data visualization libraries to display metrics in an understandable and visually pleasing manner.
  - Graphs: Time-series graphs for tracking changes over time.
  - Maps: Geographical maps to show location-based data.
  - Gauges: For real-time metrics like air and water quality.

#### Security and Compliance

- Ensure all data is securely transmitted and stored.
- Comply with GDPR or similar data protection regulations for handling user data.

#### Documentation and Help

- Provide a comprehensive FAQ section.
- Offer a tutorial or walkthrough for first-time users.

#### Future Enhancements

- Social Media Integration: Allow users to share their environmental metrics on social media.
- Community Initiatives: Provide information on local environmental events or community gardens.

---

## Earth Lens MVP
---

### Pitch

A simplified dashboard focusing on real-time environmental metrics for a specific area.

### MVP Requirements

#### Data Gathering

- Utilize at least two APIs to gather information on the following key metrics:
  - Air Quality: Retrieve data like PM2.5 and NO2 levels.
  - Water Quality: Access data on basic pollutants like heavy metals.

#### Real-Time Updates

- Use Web Sockets or Server Sent Events to update:
  - API information: Refresh environmental metrics every 5 minutes.

#### User Experience

- Geolocation: Allow users to either manually enter their location or use browser-based geolocation to get local data.
- Basic Dashboard: A simplified dashboard that shows the two key metrics (Air and Water Quality).

#### Data Visualization

- Basic Graphs: Use simple time-series graphs to show changes in Air and Water Quality over time.

#### Security

- Secure API Calls: Ensure that data transmission from the APIs is secure.

---

## Tech Stack
---

Front End (Web):

- React

- State Management
  - Nanostores

- Testing
  - Jest
  - Enzyme

- Storage
  - Local Storage
  - Cookies
  - Session Storage

- Documentation
  - JSDoc

Front End (Mobile):

- React Native

- State Management
  - Nanostores

- Testing
  - Jest
  - Enzyme

- Storage
  - React Native Secure Storage

- Documentation
  - JSDoc

Back End:

- Golang

- Routing
  - Chi Router

- Methods
  - Web Sockets 

- Testing
  - Built in Go Testing

- Documentation
  - GoDoc (documentation)


## Getting Started
---

### Step One - Build Out the Shared Back End

This project will be making use of one backend between both React Native (mobile), and React (web). This means the first step is to plan out what features we need on the backend, and start implementing essentials.

### Step Two - Build Out the React Side (Web)

I don't particularly have experience working with React Native, so building out a functional webpage to base the mobile app off of is going to be the next step.

### Step Three - Build Out the React Native Side (Mobile)

Finally work on the mobile app.


## What's in a Backend?

Well, we're going to building this out using Chi Router in Golang, and connecting it to a bunch of API endpoints. So for starters, we're going to need to build the backbone of the backend here, so we can easily connect API endpoints in.

We want this to be extensible, so let's define a structured "APIHandler" in which we can queue, de-queue, and finally connect the API endpoints from.

This will require a route structure, a source for the API key and API URL, and, if necessary, queue priority.

