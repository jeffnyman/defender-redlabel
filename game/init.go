package game

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/event"

	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/graphics"
	"github.com/jeffnyman/defender-redlabel/sound"
	"github.com/jeffnyman/defender-redlabel/state/baiter"
	"github.com/jeffnyman/defender-redlabel/state/bomber"
	"github.com/jeffnyman/defender-redlabel/state/gamestate"
	"github.com/jeffnyman/defender-redlabel/state/human"
	"github.com/jeffnyman/defender-redlabel/state/lander"
	"github.com/jeffnyman/defender-redlabel/state/player"
	"github.com/jeffnyman/defender-redlabel/state/pod"
	"github.com/jeffnyman/defender-redlabel/state/swarmer"
	"github.com/jeffnyman/defender-redlabel/systems"
	"github.com/jeffnyman/defender-redlabel/types"

	"github.com/hajimehoshi/ebiten/v2"
)

var blankImg *ebiten.Image

func (e *Engine) Init() {
	graphics.Load()

	blankImg = ebiten.NewImage(20, 20)
	blankImg.Fill(color.White)

	e.initSystems()
	e.addGame()
	e.addPlayer()
	e.initBulletPool()
	e.initBombPool()
	e.initLaserPool()
	e.initEvents()
	e.initHUD()

	defs.ScoreCharId = e.AddString("       0", 50, 80)
}

func (e *Engine) initSystems() {
	e.AddSystem(systems.NewPosSystem(true, e), UPDATE)
	e.AddSystem(systems.NewAISystem(true, e), UPDATE)
	e.AddSystem(systems.NewLifeSystem(true, e), UPDATE)
	e.AddSystem(systems.NewCollideSystem(true, e), UPDATE)
	e.AddSystem(systems.NewDrawSystem(true, e), DRAW)
	e.AddSystem(systems.NewAnimateSystem(true, e), UPDATE)
	e.AddSystem(systems.NewRadarDrawSystem(true, e), DRAW)
	e.AddSystem(systems.NewLaserDrawSystem(true, e), DRAW)
	e.AddSystem(systems.NewLaserMoveSystem(true, e), UPDATE)
}

func (e *Engine) InitEnemies() {
	for i := 0; i < defs.CurrentLevel().LanderCount; i++ {
		e.addLander(i)
	}

	for i := 0; i < defs.CurrentLevel().BaiterCount; i++ {
		e.addBaiter(i)
	}

	for i := 0; i < defs.CurrentLevel().BomberCount; i++ {
		e.addBomber(i)
	}

	for i := 0; i < defs.CurrentLevel().PodCount; i++ {
		e.addPod(i)
	}
}

func (e *Engine) InitHumans() {
	for i := 0; i < defs.CurrentLevel().HumanCount; i++ {
		e.addHuman(i)
	}
}

func (e *Engine) addGame() {
	gme := NewEntity(e, types.Player)
	gme.SetActive(true)

	sgraph := systems.NewStateGraph()
	sgraph.AddState(gamestate.NewGameIntro())
	sgraph.AddState(gamestate.NewGameStart())
	sgraph.AddState(gamestate.NewGamePlay())
	sgraph.AddState(gamestate.NewGameLevelEnd())
	sgraph.AddState(gamestate.NewGameOver())

	fsm := systems.NewFSM(sgraph)
	ai := components.NewAI(fsm, types.GameIntro)
	gme.AddComponent(ai)

}

func (e *Engine) addPlayer() {
	ssheet := graphics.GetSpriteSheet()
	plEnt := NewEntity(e, types.Player)
	defs.PlayerID = plEnt.Id
	plEnt.SetActive(true)

	x := float64(defs.WorldWidth) / 2
	y := float64(defs.ScreenHeight) / 2

	sgraph := systems.NewStateGraph()
	sgraph.AddState(player.NewPlayerPlay())
	sgraph.AddState(player.NewPlayerDie())

	fsm := systems.NewFSM(sgraph)
	ai := components.NewAI(fsm, types.PlayerPlay)
	plEnt.AddComponent(ai)
	smap := graphics.GetSpriteMap("ship.png")
	dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	plEnt.AddComponent(dr)
	col := types.ColorF{R: 1, G: 1, B: 1, A: 1}
	rd := components.NewRadarDraw(blankImg, col)
	plEnt.AddComponent(rd)
	sc := components.NewShip(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	plEnt.AddComponent(sc)
	pc := components.NewPos(x, y, 0, 0)
	plEnt.AddComponent(pc)

	fEnt := NewEntity(e, types.Player)
	fEnt.SetActive(true)
	fsmap := graphics.GetSpriteMap("thrust.png")
	fdr := components.NewDraw(ssheet, fsmap, types.ColorF{R: 1, G: 1, B: 1})

	fdr.Scale = 0.7
	fEnt.AddComponent(fdr)
	plEnt.SetChild(fEnt.Id)
	fpc := components.NewPos(0, 0, 0, 0)
	fEnt.AddComponent(fpc)
}

func (e *Engine) addLander(count int) {
	ssheet := graphics.GetSpriteSheet()
	ent := NewEntity(e, types.Lander)
	ent.SetActive(true)

	x := rand.Float64() * defs.WorldWidth

	if count < 2 {
		x = defs.WorldWidth * 0.8
	}

	pc := components.NewPos(x, defs.ScreenTop+500*rand.Float64(), 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(lander.NewLanderWait())
	sgraph.AddState(lander.NewLanderSearch())
	sgraph.AddState(lander.NewLanderMaterialize())
	sgraph.AddState(lander.NewLanderDrop())
	sgraph.AddState(lander.NewLanderGrab())
	sgraph.AddState(lander.NewLanderMutate())
	sgraph.AddState(lander.NewLanderDie())

	fsm := systems.NewFSM(sgraph)
	ai := components.NewAI(fsm, types.LanderWait)
	ai.Wait = 60 + (count%3)*200
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("lander.png")
	dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)
	col := types.ColorF{R: 0, G: 1, B: 0, A: 1}
	rd := components.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)

}

func (e *Engine) addBaiter(count int) {
	ssheet := graphics.GetSpriteSheet()
	ent := NewEntity(e, types.Baiter)
	ent.SetActive(true)

	pc := components.NewPos(0, 0, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(baiter.NewBaiterWait())
	sgraph.AddState(baiter.NewBaiterMaterialize())
	sgraph.AddState(baiter.NewBaiterHunt())
	sgraph.AddState(baiter.NewBaiterDie())

	fsm := systems.NewFSM(sgraph)
	ai := components.NewAI(fsm, types.BaiterWait)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("baiter.png")
	dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)
	col := types.ColorF{R: 0, G: 0.5, B: 0, A: 1}
	rd := components.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	sh := components.NewShootable()
	ent.AddComponent(sh)
	cl := components.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)
}

func (e *Engine) addHuman(count int) {
	ssheet := graphics.GetSpriteSheet()
	ent := NewEntity(e, types.Human)
	ent.SetActive(true)

	x := rand.Float64() * defs.WorldWidth

	if count < 2 {
		x = rand.Float64()*defs.ScreenWidth + defs.CameraX()
	}

	pc := components.NewPos(x, 0, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(human.NewHumanWalking())
	sgraph.AddState(human.NewHumanGrabbed())
	sgraph.AddState(human.NewHumanDropping())
	sgraph.AddState(human.NewHumanRescued())
	sgraph.AddState(human.NewHumanDie())

	fsm := systems.NewFSM(sgraph)
	ai := components.NewAI(fsm, types.HumanWalking)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("human.png")
	dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	ent.AddComponent(dr)
	col := types.ColorF{R: 1, G: 0, B: 1, A: 1}
	rd := components.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	sh := components.NewShootable()
	ent.AddComponent(sh)
	cl := components.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)

}

func (e *Engine) addBomber(count int) {
	ent := NewEntity(e, types.Bomber)
	ent.SetActive(true)

	x := (rand.Float64() * defs.ScreenWidth) + defs.WorldWidth/3
	y := (rand.Float64() * defs.ScreenHeight / 2) + defs.ScreenTop + 50

	pc := components.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(bomber.NewBomberMove())
	sgraph.AddState(bomber.NewBomberDie())

	fsm := systems.NewFSM(sgraph)
	ai := components.NewAI(fsm, types.BomberMove)
	ent.AddComponent(ai)

	smap := graphics.GFXFrame{
		Frame:           graphics.SourceFrame{X: 0, Y: 0, W: 20, H: 20},
		Anim_frames:     1,
		Ticks_per_frame: 30,
	}

	dr := components.NewDraw(blankImg, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Cycle = true
	dr.Bomber = true
	dr.Scale = 1
	ent.AddComponent(dr)
	col := types.ColorF{R: 0.5, G: 0, B: 1, A: 1}
	rd := components.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	cl := components.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)
	sh := components.NewShootable()
	ent.AddComponent(sh)
}

func (e *Engine) addPod(count int) {
	ent := NewEntity(e, types.Pod)
	ent.SetActive(true)

	x := (rand.Float64() * defs.ScreenWidth) + defs.WorldWidth/2
	y := (rand.Float64() * defs.ScreenHeight / 2) + defs.ScreenTop + 50

	pc := components.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(pod.NewPodMove())
	sgraph.AddState(pod.NewPodDie())

	fsm := systems.NewFSM(sgraph)
	ai := components.NewAI(fsm, types.PodMove)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("pod.png")
	ssheet := graphics.GetSpriteSheet()
	dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Scale = 1
	ent.AddComponent(dr)
	col := types.ColorF{R: 0.5, G: 0, B: 0.5, A: 1}
	rd := components.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	cl := components.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)
	sh := components.NewShootable()
	ent.AddComponent(sh)
}

func (e *Engine) AddSwarmer(count int, x, y float64) {
	ent := NewEntity(e, types.Swarmer)
	ent.SetActive(true)

	pc := components.NewPos(x, y, 0, 0)
	ent.AddComponent(pc)
	sgraph := systems.NewStateGraph()
	sgraph.AddState(swarmer.NewSwarmerMove())
	sgraph.AddState(swarmer.NewSwarmerDie())

	fsm := systems.NewFSM(sgraph)
	ai := components.NewAI(fsm, types.SwarmerMove)
	ent.AddComponent(ai)
	smap := graphics.GetSpriteMap("swarmer.png")
	ssheet := graphics.GetSpriteSheet()
	dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Scale = 1
	ent.AddComponent(dr)
	col := types.ColorF{R: 0.7, G: 0, B: 0, A: 1}
	rd := components.NewRadarDraw(blankImg, col)
	ent.AddComponent(rd)
	cl := components.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	ent.AddComponent(cl)
	sh := components.NewShootable()
	ent.AddComponent(sh)
}

func (e *Engine) AddScoreSprite(ev event.IEvent) {
	eve := ev.GetPayload().(*Entity)
	evpc := eve.GetComponent(types.Pos).(*components.Pos)

	ent := NewEntity(e, types.Score)
	ent.SetActive(true)
	pc := components.NewPos(evpc.X-defs.CameraX(), evpc.Y, 0, 0)
	pc.ScreenCoords = true
	ent.AddComponent(pc)
	s := "500.png"

	if ev.GetType() == event.HumanLandedEvent {
		s = "250.png"
	}

	smap := graphics.GetSpriteMap(s)
	ssheet := graphics.GetSpriteSheet()
	dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
	dr.Scale = 1
	ent.AddComponent(dr)
	li := components.NewLife(60)
	ent.AddComponent(li)
}

func (e *Engine) initBulletPool() {
	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 40; i++ {
		ent := NewEntity(e, types.Bullet)
		pc := components.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("bullet.png")
		dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		ent.AddComponent(dr)
		li := components.NewLife(240)
		ent.AddComponent(li)
		cl := components.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
		ent.AddComponent(cl)
		e.bulletPool = append(e.bulletPool, ent)
	}
}

func (e *Engine) initBombPool() {
	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 20; i++ {
		ent := NewEntity(e, types.Bomb)
		pc := components.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("bomb.png")
		dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		dr.Cycle = true
		ent.AddComponent(dr)
		cl := components.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
		ent.AddComponent(cl)
		li := components.NewLife(240)
		ent.AddComponent(li)
		e.bombPool = append(e.bombPool, ent)
	}
}

func (e *Engine) initLaserPool() {
	for i := 0; i < 15; i++ {
		ent := NewEntity(e, types.Laser)
		pc := components.NewPos(0, 0, 0, 0)
		ent.AddComponent(pc)
		dr := components.NewLaserDraw()
		ent.AddComponent(dr)
		li := components.NewLife(240)
		ent.AddComponent(li)
		mv := components.NewLaserMove()
		ent.AddComponent(mv)

		e.laserPool = append(e.laserPool, ent)
	}
}

func (e *Engine) initHUD() {
	ssheet := graphics.GetSpriteSheet()

	for i := 0; i < 5; i++ {
		ent := NewEntity(e, types.PlayerLife)
		ent.SetActive(true)
		pc := components.NewPos(float64(i*50), 40, 0, 0)
		pc.ScreenCoords = true
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("shiplife.png")
		dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		dr.Scale = 0.8
		if i >= defs.PlayerLives {
			dr.Hide = true
		}
		ent.AddComponent(dr)
		e.lives = append(e.lives, ent)
	}

	for i := 0; i < 5; i++ {
		ent := NewEntity(e, types.HUDBomb)
		ent.SetActive(true)
		pc := components.NewPos(350, 40+float64(i*20), 0, 0)
		pc.ScreenCoords = true
		ent.AddComponent(pc)
		smap := graphics.GetSpriteMap("smartbomb.png")
		dr := components.NewDraw(ssheet, smap, types.ColorF{R: 1, G: 1, B: 1})
		dr.Scale = 0.8
		if i >= defs.SmartBombs {
			dr.Hide = true
		}
		ent.AddComponent(dr)
		e.bombs = append(e.bombs, ent)
	}
}

func (e *Engine) initEvents() {
	start := func(ev event.IEvent) {
		e.GetSystem(types.PosSystem).SetActive(true)
		sound.Play(sound.Background)
		sound.Play(sound.Levelstart)
		e.SetPauseAll(false, -1)
		e.UpdateHUD()
	}

	playerCollide := func(ev event.IEvent) {
		en := ev.GetPayload().(*Entity)

		// logger.Info("Collide : %s ", en.GetClass().String())

		if en.GetClass() == types.Human {
			ai := en.GetComponent(types.AI).(*components.AI)

			if ai.State == types.HumanDropping {
				ai.NextState = types.HumanRescued
			}
		} else {
			e.Kill(en)
			e.SetPauseAll(true, en.GetID())
			ev := event.NewPlayerDie(e.GetPlayer())
			event.NotifyEvent(ev)
		}
	}

	explodeTrigger := func(ev event.IEvent) {
		if ct := ev.GetPayload().(*components.Pos); ct != nil {
			e.TriggerPS(ct.X, ct.Y)
		}
	}

	bulletTrigger := func(ev event.IEvent) {
		if ct := ev.GetPayload().(*components.Pos); ct != nil {
			if math.Abs(ct.DX) < 20 {
				e.TriggerBullet(ct.X, ct.Y, ct.DX, ct.DY)
				sound.Play(sound.Bullet)
			}
		}
	}

	Materialize := func(ev event.IEvent) {
		sound.PlayIfNot(sound.Materialize)
	}

	landerDie := func(ev event.IEvent) {
		defs.Score += 150
		e.ChangeString(defs.ScoreCharId, fmt.Sprintf("%8d", defs.Score))
		defs.LandersKilled++
		sound.Stop(sound.Laser)
		sound.Play(sound.Landerdie)
	}

	mutantSound := func(ev event.IEvent) {
		sound.PlayIfNot(sound.Mutant)
	}

	humanDie := func(ev event.IEvent) {
		defs.HumansKilled++

		if defs.HumansKilled == defs.CurrentLevel().HumanCount {
			e.ExplodeWorld()
			e.SetFlash(30)
			e.MutateAll()
		}

		sound.Play(sound.Humandie)
	}

	bomberDie := func(ev event.IEvent) {
		defs.Score += 250
		e.ChangeString(defs.ScoreCharId, fmt.Sprintf("%8d", defs.Score))
		sound.Stop(sound.Laser)
		sound.Play(sound.Bomberdie)
	}

	playerDie := func(ev event.IEvent) {
		pe := e.GetEntities()[defs.PlayerID]
		pai := pe.GetComponent(types.AI).(*components.AI)
		pai.NextState = types.PlayerDie
	}

	playerExplode := func(ev event.IEvent) {
		sound.Play(sound.Die)
	}

	playerFire := func(ev event.IEvent) {
		pe := ev.GetPayload().(*Entity)
		pc := pe.GetComponent(types.Pos).(*components.Pos)
		sc := pe.GetComponent(types.Ship).(*components.Ship)
		x := pc.X + 25
		y := pc.Y + 25

		if sc.Direction < 0 {
			x = pc.X - 100
		}

		e.TriggerLaser(x, y, sc.Direction)
		sound.Play(sound.Laser)
	}

	smartBomb := func(ev event.IEvent) {
		e.SetFlash(1)
		e.SmartBomb()
		e.UpdateHUD()
	}

	podDie := func(ev event.IEvent) {
		defs.Score += 1000
		e.ChangeString(defs.ScoreCharId, fmt.Sprintf("%8d", defs.Score))
		ent := ev.GetPayload().(*Entity)
		pc := ent.GetComponent(types.Pos).(*components.Pos)

		for i := 0; i < defs.CurrentLevel().SwarmerCount; i++ {
			e.AddSwarmer(i, pc.X, pc.Y)
		}

		sound.Stop(sound.Laser)
		sound.Play(sound.Poddie)
	}

	swarmerDie := func(ev event.IEvent) {
		defs.Score += 150
		e.ChangeString(defs.ScoreCharId, fmt.Sprintf("%8d", defs.Score))
	}

	baiterDie := func(ev event.IEvent) {
		defs.Score += 200
		e.ChangeString(defs.ScoreCharId, fmt.Sprintf("%8d", defs.Score))
		sound.Play(sound.Baiterdie)
	}

	HumanDropped := func(ev event.IEvent) {
		sound.Play(sound.Dropping)
	}

	humanGrabbed := func(ev event.IEvent) {
		sound.Play(sound.Grabbed)
	}

	humanRescued := func(ev event.IEvent) {
		e.AddScoreSprite(ev)
		defs.Score += 500
		e.ChangeString(defs.ScoreCharId, fmt.Sprintf("%8d", defs.Score))
	}

	humanSaved := func(ev event.IEvent) {
		e.AddScoreSprite(ev)
		defs.Score += 500
		e.ChangeString(defs.ScoreCharId, fmt.Sprintf("%8d", defs.Score))
		sound.Play(sound.Placehuman)
	}

	humanLanded := func(ev event.IEvent) {
		e.AddScoreSprite(ev)
		defs.Score += 250
		e.ChangeString(defs.ScoreCharId, fmt.Sprintf("%8d", defs.Score))
	}

	thrustOn := func(ev event.IEvent) {
		sound.Play(sound.Thruster)
	}

	thrustOff := func(ev event.IEvent) {
		sound.Stop(sound.Thruster)
	}

	event.AddEventListener(event.ExplodeEvent, explodeTrigger)
	event.AddEventListener(event.FireBulletEvent, bulletTrigger)
	event.AddEventListener(event.LanderDieEvent, landerDie)
	event.AddEventListener(event.HumanDieEvent, humanDie)
	event.AddEventListener(event.BomberDieEvent, bomberDie)
	event.AddEventListener(event.PlayerDieEvent, playerDie)
	event.AddEventListener(event.PlayerExplodeEvent, playerExplode)
	event.AddEventListener(event.StartEvent, start)
	event.AddEventListener(event.PlayerFireEvent, playerFire)
	event.AddEventListener(event.SmartBombEvent, smartBomb)
	event.AddEventListener(event.PlayerCollideEvent, playerCollide)
	event.AddEventListener(event.PodDieEvent, podDie)
	event.AddEventListener(event.BaiterDieEvent, baiterDie)
	event.AddEventListener(event.SwarmerDieEvent, swarmerDie)
	event.AddEventListener(event.HumanRescuedEvent, humanRescued)
	event.AddEventListener(event.HumanSavedEvent, humanSaved)
	event.AddEventListener(event.HumanLandedEvent, humanLanded)
	event.AddEventListener(event.HumanGrabbedEvent, humanGrabbed)
	event.AddEventListener(event.HumanDroppedEvent, HumanDropped)
	event.AddEventListener(event.PlayerThrustEvent, thrustOn)
	event.AddEventListener(event.PlayerStopThrustEvent, thrustOff)
	event.AddEventListener(event.MaterializeEvent, Materialize)
	event.AddEventListener(event.MutantSoundEvent, mutantSound)
}
