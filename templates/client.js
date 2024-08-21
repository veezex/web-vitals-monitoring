import {
  onCLS,
  onINP,
  onLCP,
} from 'https://unpkg.com/web-vitals@4/dist/web-vitals.attribution.js?module';

onCLS(metric => {
  reportMetric({
    ...sharedMetricsProps(metric),
    target: metric.attribution?.largestShiftTarget ?? '',
  });
});

onLCP(metric => {
  reportMetric({
    ...sharedMetricsProps(metric),
    target: metric.attribution.element,
  });
});

onINP(metric => {
  reportMetric({
    ...sharedMetricsProps(metric),
    target: metric.attribution.interactionTarget,
  });
});

function sharedMetricsProps(metric) {
  return {
    id: metric.id,
    name: metric.name,
    value: metric.value,
    rating: metric.rating,
    uri: location.pathname,
    client: getDeviceType(),
  }
}

function getDeviceType() {
  const userAgent = navigator.userAgent || navigator.vendor
  const screenWidth = window.innerWidth

  const mobileRegex =
    /android|avantgo|blackberry|bolt|boost|cricket|docomo|fone|hiptop|mini|mobi|palm|phone|pie|tablet|up\.browser|up\.link|webos|wos/i

  if (mobileRegex.test(userAgent) || screenWidth <= 800) {
    return 'mobile'
  }
  return 'desktop'
}

function reportMetric(value) {
  fetch('{{ .Protocol }}://{{ .Domain }}:{{ .Port }}/metric', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(value),
  });
}
