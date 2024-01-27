import FullCalendar from '@fullcalendar/react';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import listPlugin from '@fullcalendar/list';
import jaLocale from '@fullcalendar/core/locales/ja';
import SideMenu from './components/SideMenu';
function Home() {

  return (
    <div className="calendar-wrap">
      <SideMenu />
      <div className="calendar">
        <FullCalendar
          plugins={[dayGridPlugin, timeGridPlugin, listPlugin]}
          initialView="dayGridMonth"
          locales={[jaLocale]}
          locale='ja'
          headerToolbar={{
            left: 'prev,next today',
            center: 'title',
            right: 'dayGridMonth,timeGridWeek listWeek',
          }}
          events={[
            { title: 'eventを', start: '2022-03-14' },
            { title: 'こんな感じで追加できます', start: '2022-03-15', end: '2022-03-17' }
          ]}
        />
      </div>
    </div>
  );
}

export default Home;
